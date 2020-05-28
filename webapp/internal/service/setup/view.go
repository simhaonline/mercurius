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
	. "github.com/golangee/forms"
	"github.com/worldiety/mercurius/webapp/internal/client"
	"strconv"
)

const Path = "/setup"

type ContentView struct {
	*VStack
	statusBox *VStack
	btn       *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("first time setup").Style(Font(Headline2)),
		NewVStack().Self(&view.statusBox),


		NewButton("check").AddClickListener(func(v View) {
			go func() {
				client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
					view.statusBox.RemoveAll()
					for _,status := range res{
						view.statusBox.AddViews(NewText("*"+strconv.Itoa(status.Id)+":"+status.Message).Style(Font(Body)))
					}
				})
			}()


		}),


	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
