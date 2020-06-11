package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewLicenseView() View {
	values := NewResources(locale.Language())

	return NewVStack(
		NewText(values.SetupLicense()),
	).SetHorizontalAlign(Center)
}
