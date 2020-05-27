// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package setup

import (
	"net/http"
)

type Reloader interface {
	ReloadStatus() []error
	Reload()
}

// Status represents the current setup status.
type Status struct {
	// status id
	Id int
	// a textual representation as a developer notice
	Message string
}

// @ee.http.Controller
// @ee.http.Route("/api/v1/setup")
// @ee.stereotype.Controller("setup")
type RestController struct {
	ctr Reloader
}

func NewRestController(ctr Reloader) *RestController {
	return &RestController{ctr}
}

// Status returns the current setup status. This is usually only relevant in the installation phase.
//
// @ee.http.Route("/status")
// @ee.http.Method("GET")
func (s *RestController) Status(res http.ResponseWriter, req *http.Request) []Status {
	return []Status{{Id: -1, Message: "hello world"}}
}