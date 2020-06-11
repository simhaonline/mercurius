package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewWelcomeView() View {
	values := NewResources(locale.Language())

	return NewVStack(
		NewText(values.SetupWelcome()),
	).SetHorizontalAlign(Center)
}
