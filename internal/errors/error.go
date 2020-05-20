// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

// +build !js,!wasm

package errors

// Try executes the given func and updates the error,
// but only if it has not been set yet.
func Try(f func() error, err *error) {
	newErr := f()

	if *err == nil {
		*err = newErr
	}
}
