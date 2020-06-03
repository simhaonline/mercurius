//go:generate go run main.go
package main

import (
	"github.com/golangee/i18n"
)

func main() {
	// invoke the generator in your current project. It will process the entire module.
	if err := i18n.Bundle(); err != nil {
		panic(err)
	}
}
