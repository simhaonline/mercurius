package main

import (
	"github.com/worldiety/gluon"
	"github.com/worldiety/mercurius/specification/resources/user"
)

func Define() *gluon.Package {
	p := gluon.NewPackage("mercurius", nil)
	user.Define(p)
	return p
}

func main() {
	diagram, err := gluon.ClassDiagram(Define(), "svg")
	if err != nil {
		panic(err)
	}
	err = diagram.WriteFile("/Users/tschinke/tmp/mercurius-classDiagram.svg")
	if err != nil {
		panic(err)
	}
}
