//go:generate go run gen.go
package main

import (
	"github.com/golangee/reflectplus"
)

func main() {
	reflectplus.Must(reflectplus.Generate("../.."))
}
