// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package notfound

import (
	. "github.com/golangee/forms"
)

const Path = "/demo"

type ContentView struct {
	*VStack
}

func NewContentView(path string) *ContentView {
	return &ContentView{VStack: NewVStack().AddViews(
		NewText("the route '"+path+"' is not available").Style(Font(Headline1)),
		NewButton("Index").AddClickListener(func(v View) {
			v.Context().Navigate("/")
		}),
	)}
}

func FromQuery(q Query) View {
	return NewContentView(q.Path())
}
