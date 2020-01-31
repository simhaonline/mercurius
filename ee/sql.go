/*
 * Copyright 2020 Torben Schinke
 *
 * worldiety Enterprise Edition (EE) License
 * See the file LICENSE for allowed usage and redistribution information.
 *
 * Please contact worldiety GmbH or visit www.worldiety.de if you need additional information or have any
 * questions.
 */

package ee

import (
	"database/sql"
	"fmt"
)

// ErrNoDatabaseAvailable is returned, if no transaction can be opened.
var ErrNoDatabaseAvailable = fmt.Errorf("no database available")

// SQL holds a valid transaction
type SQL struct {
	parent        *Ctx    // parent is never nil
	tx            *sql.Tx // tx may only be nil, if firstError is not nil
	firstError    error   // may be nil
	pendingResult *Result // pendingResult contains the currently pending result.
}

// NewSQL tries to create a new sql transaction. If that fails, either because there is no db in the context
// or because a transaction cannot be opened, the SQL instance is set to failed but can be used in queries.
// However, there will be never a real sql call anymore.
func NewSQL(ctx *Ctx) *SQL {
	s := &SQL{parent: ctx}
	if ctx.db == nil {
		s.firstError = ErrNoDatabaseAvailable
	} else {
		tx, err := ctx.db.BeginTx(ctx.context, ctx.opts)
		if err != nil {
			s.firstError = fmt.Errorf("cannot begin transaction: %w", err)
		}
		s.tx = tx
	}

	return s
}

// Tx returns the currently wrapped sql transaction or nil if Error is set
func (s *SQL) Tx() *sql.Tx {
	return s.tx
}

// Error returns the first occurred error.
func (s *SQL) Error() error {
	return s.firstError
}

// Query prepares a query and delegates to a context. Never fails.
func (s *SQL) Query(query string, params ...interface{}) *Result {
	if s.firstError != nil {
		return &Result{parent: s}
	}

	if s.tx == nil {
		panic("tx must not be nil without error")
	}

	stmt, err := s.tx.PrepareContext(s.parent.Context(), query)
	if err != nil {
		s.firstError = fmt.Errorf("unable to prepare statement: %s [%v]:%w", query, params, err)
		return &Result{parent: s}
	}

	rows, err := stmt.Query(params...) // nolint: rowserrcheck // we ensure that close and err are evaluated later
	if err != nil {
		s.firstError = fmt.Errorf("unable to execute statement: %s [%v]:%w", query, params, err)
		return &Result{parent: s}
	}

	res := &Result{parent: s, rows: rows}
	_ = res.parent.closePending()
	s.pendingResult = res

	return res
}

// close tries to close any pending result
func (s *SQL) closePending() error {
	if s.pendingResult != nil {
		err := s.pendingResult.rows.Close()

		if s.firstError != nil {
			s.firstError = fmt.Errorf("failed to close pending rows: %w", err)
		}

		s.pendingResult = nil

		return err
	}

	return s.firstError
}

// Close tries to commit if rollback is false and no error is pending. Safe to call on a nil pointer, which is not an
// error.
func (s *SQL) Close(rollback bool) error {
	if s == nil {
		return nil
	}

	if s.tx != nil {
		if s.firstError != nil || rollback {
			err := s.tx.Rollback()

			if s.firstError != nil {
				s.firstError = err
			}
		} else {
			s.firstError = s.tx.Commit()
		}
	}

	return s.firstError
}

type Result struct {
	parent *SQL      // parent is never nil
	rows   *sql.Rows // result may be nil. parent.Error tells you why
}

type Row interface {
	Scan(dst ...interface{}) error
}

// Map calls the given closure for each row of the result and returns the first error.
func (r *Result) Map(closure func(row Row) error) error {
	if r.parent.firstError != nil {
		return r.parent.firstError
	}

	if r.rows == nil {
		panic("rows must not be nil without error")
	}

	for r.rows.Next() {
		err := closure(r.rows)
		if err != nil {
			r.parent.firstError = fmt.Errorf("failed to map: %w", err)
			r.parent.parent.Suppressed(r.rows.Close())
			r.parent.pendingResult = nil

			return r.parent.firstError
		}
	}

	if r.rows.Err() != nil {
		r.parent.firstError = fmt.Errorf("failed to iterate: %w", r.rows.Err())
		r.parent.parent.Suppressed(r.rows.Close())
		r.parent.pendingResult = nil

		return r.parent.firstError
	}

	err := r.rows.Close()
	if err != nil {
		r.parent.firstError = fmt.Errorf("failed to close rows: %w", r.rows.Err())
		r.parent.pendingResult = nil

		return r.parent.firstError
	}

	r.parent.pendingResult = nil

	return nil
}
