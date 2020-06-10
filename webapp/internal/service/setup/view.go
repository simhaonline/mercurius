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
	tabView      *TabView
	content      *Frame
	currentIndex int
	stepper      *Stepper
	title        *Text
	strings      Resources
	btnNext      *Button
}

func NewContentView() *ContentView {

	view := &ContentView{}
	view.VStack = NewVStack()
	view.currentIndex = 1
	view.strings = NewResources(locale.Language())

	client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
		if client.FindError(err, errors.MercuriusConfigurationMissing) == nil {
			view.VStack.ClearViews()
			view.VStack.AddViews(NewText("nothing to do"))
			return
		}

		view.VStack.AddViews(
			NewHStack(NewText("header").
				Self(&view.title).
				Style(Font(Headline2), Margin()),
			).SetHorizontalAlign(Center),
			NewStepper(
				NewIconStep(icon.AirlineSeatFlat, "welcome"),
				NewIconStep(icon.Assignment, "License"),
				NewIconStep(icon.Storage, "Database"),
				NewIconStep(icon.Folder, "File storage"),
				NewIconStep(icon.Settings, "http"),
			).Self(&view.stepper).SetProgress(view.currentIndex),
			NewFrame().Self(&view.content),
			NewHStack(
				NewButton("next").
					Self(&view.btnNext).
					SetStyleKind(Raised).
					AddClickListener(func(v View) {
						view.NextAction()
					}),
			).Style(Padding()).SetHorizontalAlign(Center),
		).Style(Height(Percent(100))).SetRowHeights(Auto(), Auto(), Fraction(1), Auto())

		view.ShowWelcome()
	})

	return view
}

func (t *ContentView) NextAction() {
	t.currentIndex++
	t.stepper.SetProgress(t.currentIndex)
	switch t.currentIndex {
	case 2:

	}
}

func (t *ContentView) ShowWelcome() {
	t.title.Set(t.strings.SetupTitleWelcome())
	t.content.SetView(NewWelcomeView())
	t.btnNext.SetText(t.strings.BtnAccept())
}

func FromQuery(Query) View {
	return NewContentView()
}
