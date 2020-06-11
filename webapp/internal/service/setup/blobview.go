package setup

import (
	. "github.com/golangee/forms"
	"github.com/golangee/forms/locale"
)

func NewBlobView(cfg *BlobStore) View {
	values := NewResources(locale.Language())

	return NewVStack(
		NewPicker("filesystem").
			SetSelected(0).
			SetLabel(values.SetupDriver()).
			Style(Width(Percent(100))),
		NewTextField().
			SetText("~/.mercurius/storage").
			SetLabel(values.SetupPath()).
			Style(Width(Percent(100))),
	)
}
