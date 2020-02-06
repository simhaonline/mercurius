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
	"context"
	"database/sql"
	"fmt"
)

// Ctx represents a typed context object for any request (or other scoped operation) issued in an application.
// It bundles the Context to use for timeouts, deadlines or cancellations and a database transaction. It is unique
// per job and goroutine and never shared.
//
// You must never keep a reference to this Requests nor to anything it contains or returns.
type Ctx struct {
	context context.Context // context is never nil

	opts    *sql.TxOptions  // opts may be nil

	// multiple db connections
	sql     *SQL            // sql is always nil at first
	db *sql.DB // db may be nil, if no db configured
	// Logger?
	// Session? -> no not always an http request
}

func NewCtx(ctx context.Context, db *sql.DB) *Ctx {
	return &Ctx{context: ctx, db: db}
}

// Context returns the go context with timeout, deadline and cancellation to use.
func (c *Ctx) Context() context.Context {
	return c.context
}

func (c *Ctx) Tx() (*sql.Tx, error) {
	return nil, nil
}

// SQL returns the wrapped database transaction to use. It will allocate everything in a lazy way and commits everything
// on execution. It will never fail and will never return nil.
func (c *Ctx) SQL() *SQL {
	if c.sql == nil {
		c.sql = NewSQL(c)
	}

	return c.sql
}

// reset just nils the sql state, so that it can get GC'ed and the context can be reused
func (c *Ctx) reset() {
	c.sql = nil
}

// Exec guarantees that a panic can never happen and that any pending transaction is either rolled back (if any error
// occurs) or else committed. The Ctx can be reused after this call.
func (c *Ctx) Exec(f func() error) error {
	defer c.reset()

	err := WithPanic("Ctx.Exec", f)
	if err != nil {
		c.Suppressed(c.sql.Close(true))
		return err
	}

	return c.sql.Close(false)
}

// Suppressed can be used to note aftereffect errors, which otherwise would become unnoticed.
func (c *Ctx) Suppressed(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
