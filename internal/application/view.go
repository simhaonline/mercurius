// +build js,wasm

package application

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/theme/material/icon"
	"github.com/worldiety/mercurius/build"
	"path/filepath"
	"time"
)

type App struct {
	*Application
}

func NewApp() *App {
	a := &App{}
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
