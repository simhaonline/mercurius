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

import "runtime/debug"

// The ErrPanic
type ErrPanic struct {
	Cause   interface{} // Error contains the recovered error from the panic
	Stack   string      // Stack contains a formatted string of the failed trace
	Message string      // Message contains a context specific message
}

// Error returns the message and the stack
func (e ErrPanic) Error() string {
	return e.Message + ": recovered panic:\n" + e.Stack
}

// Unwrap returns the error cause, if it can be casted as an error. Otherwise you should inspect Cause directly.
func (e ErrPanic) Unwrap() error {
	if err, ok := e.Cause.(error); ok {
		return err
	}

	return nil
}

// WithPanic executes the given closure and tries to call it. Any panic is wrapped within an ErrPanic, otherwise
// the error value from f is just returned. Thus it is guaranteed that a call on WithPanic never panics.
func WithPanic(msg string, f func() error) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = ErrPanic{
				Cause:   p,
				Stack:   string(debug.Stack()),
				Message: msg,
			}
		}
	}()

	err = f()

	return
}

// Must consumes an error and panics if err is not nil. Use this only to make assertions and to indicate a programming
// error and not a regular error condition.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
