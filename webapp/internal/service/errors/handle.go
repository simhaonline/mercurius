package errors

import (
	. "github.com/golangee/forms"
	"github.com/worldiety/mercurius/webapp/internal/client"
)

const (
	MercuriusConfigurationMissing = "hg.configuration.missing"
)

// HandleError checks the error and shows e.g. an Error dialog and may perform even a redirect etc.
func HandleError(target ViewGroup, err error) bool {
	if err != nil {
		target.ClearViews()
		if client.FindError(err, MercuriusConfigurationMissing) != nil {
			target.Context().Navigate("/setup")
		}
		target.AppendViews(NewErrorView(err))
		return true
	}
	return false
}
