package errors

import (
	. "github.com/golangee/forms"
)

// HandleError checks the error and shows e.g. an Error dialog and may perform even a redirect etc.
func HandleError(target ViewGroup, err error) bool {
	if err != nil {
		target.ClearViews()
		target.AppendViews(NewErrorView(err))
		return true
	}
	return false
}

