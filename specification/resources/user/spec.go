package user

import "github.com/worldiety/gluon"

func Define(parent *gluon.Package) {
	gluon.NewClass("User", parent).
		AddFields(
			gluon.NewField(false, "id"),
			gluon.NewField(true, "fkId"),
		)
}
