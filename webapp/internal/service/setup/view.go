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
	"github.com/golangee/forms/locale"
	"github.com/worldiety/mercurius/webapp/internal/client"
	"github.com/worldiety/mercurius/webapp/internal/service/errors"
)

const Path = "/setup"

type ContentView struct {
	*VStack
}

func NewContentView() *ContentView {
	values := NewResources(locale.Language())

	view := &ContentView{}
	view.VStack = NewVStack().AddViews(

		NewText("setup").Style(Font(Headline3)),


	)

	client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
		if client.FindError(err, errors.MercuriusConfigurationMissing) == nil {
			view.VStack.ClearViews()
			view.VStack.AddViews(NewText("nothing to do"))
			return
		}

		view.VStack.AddViews(
			NewTabView().SetTabs(
				NewTab(values.SetupTitleLicense(), NewText(values.SetupLicense())),
			),
		)
	})

	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
