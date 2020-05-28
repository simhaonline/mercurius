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
	"io/ioutil"
	"net/http"
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
				//client.NewMercuriusService("http://localhost:8080/","blub",nil)

				res, err := http.Get("/api/v1/setup/status")
				if err != nil {
					ShowMessage(v, err.Error())
					return
				}
				defer res.Body.Close()
				b, err := ioutil.ReadAll(res.Body)
				if err != nil {
					ShowMessage(v, err.Error())
					return
				}
				view.statusBox.RemoveAll()
				view.statusBox.AddViews(
					NewText(string(b)).Style(Font(Body)),
				)

			}()


		}),


	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
