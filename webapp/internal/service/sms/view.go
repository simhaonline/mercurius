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
	. "github.com/golangee/forms"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("sms").Style(Font(Headline1)),
		NewText("need to do the setup").Style(Font(Body)),

	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
