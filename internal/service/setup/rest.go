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
	"fmt"
	"github.com/golangee/uuid"
	"io"
	"net/http"
	"time"
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
	ID2     uuid.UUID
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
// @ee.http.Route("/status2/:id")
// @ee.http.PathParam("id")
// @ee.http.QueryParam("x")
// @ee.http.Method("GET")
func (s *RestController) Status2(res http.ResponseWriter, id,x uuid.UUID) ([]Status, error) {
	return nil,fmt.Errorf("blub")
}

// Status returns the current setup status. This is usually only relevant in the installation phase.
//
// @ee.http.Route("/status")
// @ee.http.Method("GET")
func (s *RestController) Status(res http.ResponseWriter, req *http.Request) ([]Status, error) {
	var r []Status
	for _, err := range s.ctr.ReloadStatus() {
		r = append(r, Status{
			Id:      0,
			Message: err.Error(),
		})
	}
	time.Sleep(2 * time.Second)
	return r, fmt.Errorf("failed test badly: %w", io.EOF)
}
