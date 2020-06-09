// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

package application

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"github.com/golangee/log"
	"github.com/worldiety/mercurius/webapp/build"
	"github.com/worldiety/mercurius/webapp/internal/service/dashboard"
	"github.com/worldiety/mercurius/webapp/internal/service/notfound"
	"github.com/worldiety/mercurius/webapp/internal/service/setup"
	"path/filepath"
	"time"
)

type App struct {
	*Application
}

func NewApp() *App {
	a := &App{}
	logger := log.New("")
	logger.Info("mercurius frontend", log.Obj("version", build.Env().String()))
	a.Application = NewApplication(a, build.Env().String())
	return a
}

func (a *App) WithDrawer(f func(q Query) View) func(Query) View {
	return func(query Query) View {
		v := NewGroup(f(query)).Style(Padding())

		var items []LstItem
		items = append(items,
			NewListItem("home").SetLeadingView(NewIcon(icon.Home)).AddClickListener(func(v View) {
				a.Context().Navigate("/")
			}),
			NewListSeparator(),
			NewListHeader("components"),
		)
		for _, route := range a.Context().Routes() {
			fPath := route.Path
			name := filepath.Base(route.Path)
			if fPath == "/" {
				continue
			}
			item := NewListItem(name)
			if route.Path == query.Path() {
				item.SetSelected(true)
			}
			items = append(items, item.AddClickListener(func(v View) {
				go func() {
					time.Sleep(200 * time.Millisecond) // wait for drawer animation
					a.Context().Navigate(fPath)
				}()
			}))
		}

		return NewDrawer(
			NewTopAppBar().
				SetTitle("wtk demo").
				SetNavigation(icon.Menu, nil).
				AddActions(NewIconItem(icon.Help, "download", func(v View) {
					ShowMessage(v, "wtk demo")
				})),
			NewVStack().AddViews(
				NewText("your demo").Style(Font(DrawerTitle)),
				NewText("anonymous").Style(Font(DrawerSubTitle)),
			),
			NewList().AddItems(items...),
			v)
	}
}


func NoDrawerFixedBox(window *Window, f func(q Query) View) func(Query) View {
	return func(query Query) View {
		v := NewVStack(
			NewCard(f(query)).
				Style(Width(Pixel(1200)), Height(Pixel(670))).
				StyleFor(MatchOne(MatchMaxWidth(Pixel(1200)), MatchMaxHeight(Pixel(670))), Width(PercentViewPortWidth(100)), Height(PercentViewPortHeight(100))),

		).Style(Height(Percent(100))).Grid.SetVerticalAlign(Center).SetHorizontalAlign(Center)
		window.SetBackground("/hg/img/background-01.jpg")
		return v
	}
}

func (a *App) Start() {
	Theme().SetColor(0x1b8c30ff)

	a.UnmatchedRoute(notfound.FromQuery)
	a.Route(setup.Path, NoDrawerFixedBox(a.Window(), setup.FromQuery))
	a.Route(dashboard.Path, a.WithDrawer(dashboard.FromQuery))
	a.Route("/", a.WithDrawer(dashboard.FromQuery))
	a.Application.Start()
}
