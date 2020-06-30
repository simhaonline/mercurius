package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewHttpView(cfg *Server) View {
	values := NewResources(locale.Language())

	return NewGrid().SetAreas([][]string{
		{"adr", "adr", "adr", "port"},
	}).SetGap(DefaultPadding).
		AddView(
			NewTextField().
				SetLabel(values.SetupHttpAddress()).
				BindText(&cfg.Address).
				Style(Width(Percent(100))),
			GridLayoutParams{Area: "adr"}).
		AddView(
			NewTextField().
				SetLabel(values.SetupPort()).
				BindInt(&cfg.Port).
				Style(Width(Percent(100))),
			GridLayoutParams{Area: "port"},
		)

}
