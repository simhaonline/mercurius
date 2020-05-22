// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package dashboard

import (
	. "github.com/golangee/forms"
)

const Path = "/dashboard"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("your dashboard").Style(Font(Headline2)),
		NewText("your account").Style(Font(Body)),

	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
