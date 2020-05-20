// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package sms

import (
	"context"
	"fmt"
	"github.com/golangee/uuid"
	"net/http"
)

// @ee.http.Controller
// @ee.http.Route("/api/v1/sms")
type RestController struct {
	sms MessageRepository
}

func NewRestController(sms MessageRepository) *RestController {
	return &RestController{sms}
}

// @ee.http.QueryParam("limit")
// @ee.http.Method("GET")
func (s *RestController) List(ctx context.Context, limit int) ([]SMS, error) {
	return s.sms.FindAll(ctx, limit)
}

// @ee.http.HeaderParam("value":"userAgent","alias":"User-Agent")
// @ee.http.Route("/a/:id")
// @ee.http.Method("GET")
func (s *RestController) Get(ctx context.Context, id uuid.UUID, userAgent string) (SMS, error) {
	fmt.Println(userAgent)
	return s.sms.FindById(ctx, id)
}

// @ee.http.Route("/asdf")
// @ee.http.Method("GET")
func (s *RestController) Get2(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}
