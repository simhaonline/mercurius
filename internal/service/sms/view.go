package sms

import (
	. "github.com/golangee/forms"
)

const Path = "/demo/topappbar"

type ContentView struct {
	*VStack
	btn *Button
}

func NewContentView() *ContentView {
	view := &ContentView{}
	view.VStack = NewVStack().AddViews(
		NewText("first time setup").Style(Font(Headline1)),
		NewText("need to do the setup").Style(Font(Body)),

	)
	return view
}

func FromQuery(Query) View {
	return NewContentView()
}
