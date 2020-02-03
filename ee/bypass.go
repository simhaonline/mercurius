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
	"fmt"
	"strings"
)

// A ByPassError wraps the first error and provides a list of suppressed errors. It has no own message to add.
type ByPassError struct {
	First      error   // First contains the first error
	Suppressed []error // Suppressed contains all those errors which occurred after the first error.
}

// Add appends the given err as the first or a suppressed error. This means that the error occurred after the
// first error is suppressed. However it does not necessarily mean that it is indeed a consequential error.
func (e *ByPassError) Add(err error) {
	if err == nil {
		return
	}
	if e.First == nil {
		e.First = err
	} else {
		e.Suppressed = append(e.Suppressed, err)
	}
}

// Unwrap returns nil or First. Sadly the error contract of go does not allow to inspect suppressed errors, however
// it is likely less important anyway.
func (e *ByPassError) Unwrap() error {
	return e.First
}

// Error returns a string representation of this error. It inserts the suppressed errors before the first colon.
func (e *ByPassError) Error() string {
	if e.First == nil {
		return fmt.Sprintf("Empty ByPassError (suppressed: %d)", len(e.Suppressed))
	}

	if len(e.Suppressed) > 0 {
		sb := &strings.Builder{}
		messages := strings.Split(e.First.Error(), ":")
		if len(messages) > 0 {
			sb.WriteString(messages[0])
			sb.WriteByte(' ')
		}

		sb.WriteByte('[')
		for i, err := range e.Suppressed {
			sb.WriteString("suppressed '")
			sb.WriteString(err.Error())
			if i < len(e.Suppressed)-1 {
				sb.WriteString("', ")
			}
		}
		sb.WriteByte(']')
		sb.WriteByte(':')
		sb.WriteString(strings.Join(messages[1:], ":"))

		return sb.String()
	}
	return e.First.Error()

}

// A ByPass is a helper which allows to implement dsl-like (builder) patterns which perform internally
// a short circuit to write more fluid code. This should only be used in situations where error handling
// would obscure the actual job, especially when working with a set of related functions.
// If you encounter such a situation, you should introduce another layer of abstraction.
type ByPass struct {
	error *ByPassError
}

// Error returns either nil or a *ByPassError
func (b *ByPass) Error() error {
	return b.error
}

// Try returns true, if you should continue with whatever you are doing. Pass in your error value and the condition
// is checked and logged as needed. If any error has been inspected, Try will always return false, even if err is nil.
func (b *ByPass) Try(err error) bool {
	if err != nil {
		if b.error == nil {
			b.error = &ByPassError{}
		}
		b.error.Add(err)
	}
	return b.error == nil
}

// Fail is like Try but the inverse. So if err is not nil, false is returned. Besides the return value, they
// behave exactly the same, so use what reads best.
func (b *ByPass) Fail(err error) bool {
	return !b.Try(err)
}

// Reset removes the internal ByPass error, just as a clean new allocation.
func (b *ByPass) Reset() {
	b.error = nil
}
