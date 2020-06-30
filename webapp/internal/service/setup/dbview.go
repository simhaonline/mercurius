package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewDBView(cfg *Sql) View {
	values := NewResources(locale.Language())

	grid := NewGrid().SetAreas([][]string{
		{"driver", "driver", "driver", "driver"},
		{"host", "host", "host", "port"},
		{"name", "name", "name", "name"},
		{"user", "user", "pwd", "pwd"},
		{"ssl", "ssl", "ssl", "ssl"},
	}).SetGap(DefaultPadding)

	grid.AddView(
		NewPicker("mysql", "b", "c").
			SetSelected(0).
			BindText(&cfg.Driver).
			SetLabel(values.SetupDriver()).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "driver"})

	grid.AddView(
		NewTextField().
			SetLabel(values.SetupDbHost()).
			BindText(&cfg.Host).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "host"})

	grid.AddView(
		NewTextField().
			SetLabel(values.SetupPort()).
			SetInputType(Number).
			BindInt(&cfg.Port).
			SetText("3306").
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "port"})

	grid.AddView(
		NewTextField().
			SetLabel(values.SetupDbName()).
			BindText(&cfg.DatabaseName).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "name"})

	grid.AddView(
		NewTextField().
			SetLabel(values.SetupUser()).
			BindText(&cfg.User).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "user"})

	grid.AddView(
		NewTextField().
			SetLabel(values.SetupPassword()).
			SetInputType(Password).
			BindText(&cfg.Password).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "pwd"})

	grid.AddView(
		NewPicker(
			"Preferred",
			"Disable",
			"Required",
			"VerifyCA",
			"VerifyIdentify").
			SetSelected(0).
			BindText(&cfg.SSLMode).
			SetLabel(values.SetupSslMode()).
			Style(Width(Percent(100))),
		GridLayoutParams{Area: "ssl"})

	return grid
}
