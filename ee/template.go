package ee

import "io"

// A Template may have different implementations, e.g. a text or html template.
type Template interface {
	// Execute renders the template into the writer.
	// It is safe to be executed concurrently on different writers.
	Execute(wr io.Writer, data interface{}) error
}

// A ReloadableTemplate watches a concrete file and reloads
type ReloadableTemplate struct{
	Filename string
}

func NewReloadableTemplate(fname string)ReloadableTemplate{

}