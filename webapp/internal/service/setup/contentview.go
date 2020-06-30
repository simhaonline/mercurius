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
	btnBack      *Button
	config       *Settings
}

func NewContentView() *ContentView {

	view := &ContentView{}
	view.VStack = NewVStack()
	view.currentIndex = 1
	view.strings = NewResources(locale.Language())
	view.config = &Settings{
		BlobStore: BlobStore{},
		Database: Sql{
			Driver:       "",
			Host:         "localhost",
			Port:         3306,
			User:         "",
			Password:     "",
			DatabaseName: "",
			SSLMode:      "",
		},
		Server: Server{
			Port:    8080,
			Address: "localhost",
		},
		Development: false,
	}

	client.Service().SetupService().ApiV1SetupStatus(view.Scope(), func(res []client.Status, err error) {
		if client.FindError(err, errors.MercuriusConfigurationMissing) == nil {
			view.VStack.ClearViews()
			view.VStack.AddViews(NewText("nothing to do"))
			return
		}

		view.VStack.AddViews(
			NewHStack(NewText("").
				Self(&view.title).
				Style(Font(Headline2), Margin()),
			).SetHorizontalAlign(Center),
			NewStepper(
				NewIconStep(icon.Home, view.strings.SetupStepperWelcome()),
				NewIconStep(icon.Assignment, view.strings.SetupStepperLicense()),
				NewIconStep(icon.Storage, view.strings.SetupStepperDatabase()),
				NewIconStep(icon.Folder, view.strings.SetupStepperStorage()),
				NewIconStep(icon.Settings, view.strings.SetupStepperHttp()),
			).Self(&view.stepper).SetProgress(view.currentIndex),
			NewFrame().Self(&view.content).Style(Overflow(OverflowAuto), Padding()),
			NewHStack(
				NewButton(view.strings.BtnBack()).
					Self(&view.btnBack).
					SetStyleKind(Raised).
					AddClickListener(func(v View) {
						view.BackAction()
					}).
					Style(MarginRight(DefaultPadding)),
				NewButton("").
					Self(&view.btnNext).
					SetStyleKind(Raised).
					AddClickListener(func(v View) {
						view.NextAction()
					}),
			).Style(Padding()).SetHorizontalAlign(Center),
		).Style(Height(Percent(100))).SetRowHeights(Auto(), Auto(), Fraction(1), Auto())

		view.showWelcome()
	})

	return view
}

func (t *ContentView) NextAction() {
	t.currentIndex++
	t.applyState()
}

func (t *ContentView) BackAction() {
	t.currentIndex--
	t.applyState()
}

func (t *ContentView) applyState() {
	t.stepper.SetProgress(t.currentIndex)

	switch t.currentIndex {
	case 1:
		t.showWelcome()
	case 2:
		t.showLicense()
	case 3:
		t.showDB()
	case 4:
		t.showBlob()
	case 5:
		t.showHttp()
	}
}

func (t *ContentView) showWelcome() {
	t.title.Set(t.strings.SetupTitleWelcome())
	t.content.SetView(NewWelcomeView())
	t.btnNext.SetText(t.strings.BtnStart())
	t.btnBack.Style(Display(DisplayNone))
}

func (t *ContentView) showLicense() {
	t.title.Set(t.strings.SetupTitleLicense())
	t.content.SetView(NewLicenseView())
	t.btnNext.SetText(t.strings.BtnAccept())
	t.btnBack.Style(Display(DisplayVisible))
}

func (t *ContentView) showDB() {
	t.title.Set(t.strings.SetupTitleDb())
	t.content.SetView(NewDBView(&t.config.Database))
	t.btnNext.SetText(t.strings.BtnNext())
	t.btnBack.Style(Display(DisplayVisible))
}

func (t *ContentView) showBlob() {
	t.title.Set(t.strings.SetupTitleBlobs())
	t.content.SetView(NewBlobView(&t.config.BlobStore))
	t.btnNext.SetText(t.strings.BtnNext())
	t.btnBack.Style(Display(DisplayVisible))
}

func (t *ContentView) showHttp() {
	t.title.Set(t.strings.SetupTitleHttp())
	t.content.SetView(NewHttpView(&t.config.Server))
	t.btnNext.SetText(t.strings.BtnApply())
	t.btnBack.Style(Display(DisplayVisible))
}

func FromQuery(Query) View {
	return NewContentView()
}
