//go:generate go run gen.go
package main

import (
	"github.com/worldiety/sqlm"
)

func main() {
	sqlm.Must(sqlm.GenerateAll("../.."))
}