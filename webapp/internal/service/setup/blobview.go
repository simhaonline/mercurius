package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewBlobView(cfg *BlobStore) View {
	values := NewResources(locale.Language())

	return NewVStack(
		NewPicker("filesystem", "c").
			SetSelected(0).
			BindText(&cfg.Driver).
			SetLabel(values.SetupDriver()).
			Style(Width(Percent(100))),
		NewTextField().
			SetText("~/.mercurius/storage").
			BindText(&cfg.Path).
			SetLabel(values.SetupPath()).
			Style(Width(Percent(100))),
	)
}
