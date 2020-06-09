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
	"github.com/golangee/forms/theme/material/icon"
	. "github.com/golangee/forms/views/hstepper"
	"github.com/worldiety/mercurius/webapp/internal/client"
	"github.com/worldiety/mercurius/webapp/internal/service/errors"
)

const Path = "/setup"

type ContentView struct {
	*VStack
	tabView *TabView
}

func NewContentView() *ContentView {
	values := NewResources(locale.Language())
	_ = values
	view := &ContentView{}
	view.VStack = NewVStack()

	client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
		if client.FindError(err, errors.MercuriusConfigurationMissing) == nil {
			view.VStack.ClearViews()
			view.VStack.AddViews(NewText("nothing to do"))
			return
		}

		view.VStack.AddViews(
			NewHStack(NewText("header")).SetHorizontalAlign(Center).Style(BackgroundColor(Gray50)),
			NewStepper(
				NewIconStep(icon.Assignment, "License"),
				NewIconStep(icon.Storage, "Database"),
				NewIconStep(icon.Folder, "File storage"),
				NewIconStep(icon.Settings, "http"),
			).Style(BackgroundColor(Yellow50)).SetProgress(2),
			NewText("content").Style(BackgroundColor(Blue50)),
			NewHStack(NewButton("next").SetStyleKind(Raised)).SetHorizontalAlign(End).Style(BackgroundColor(Red50)),
		).Style(Height(Percent(100))).SetRowHeights(Auto(), Auto(), Fraction(1), Auto())
	})

	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
