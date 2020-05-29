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
	"github.com/golangee/log"
	"github.com/worldiety/mercurius/webapp/internal/client"
	"github.com/worldiety/mercurius/webapp/internal/service/errors"
	"reflect"
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


		NewButton("check2").AddClickListener(func(v View) {
			view.statusBox.AddViews(NewCircularProgress())
			client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
				log.New("setup").Info("view", log.Obj("err", err), log.Obj("nil", err == nil), log.Obj("ref", reflect.TypeOf(err)))
				errors.HandleError(view, err)
				view.statusBox.RemoveAll()
				for _, status := range res {
					view.statusBox.AddViews(NewText("*" + strconv.Itoa(status.Id) + ":" + status.Message).Style(Font(Body)))
				}

			})

		}),


	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
