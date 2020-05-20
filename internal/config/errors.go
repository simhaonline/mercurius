// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package config

// FirstTimeSetupError indicates that a setup has never been made (clean install).
type FirstTimeSetupError struct {
}

func (f FirstTimeSetupError) Error() string {
	return "first time setup required"
}

// NoDatabaseError indicates a missing database or wrong credentials
type NoDatabaseError struct {
	Cause error
}

func (f NoDatabaseError) Error() string {
	return "cannot open database"
}

func (f NoDatabaseError) Unwrap() error {
	return f.Cause
}

// InvalidConfigurationError indicates that a setup was made but is broken
type InvalidConfigurationError struct {
	Cause error
}

func (f InvalidConfigurationError) Error() string {
	return "invalid configuration"
}

func (f InvalidConfigurationError) Unwrap() error {
	return f.Cause
}
