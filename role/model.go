/*
 * Copyright 2020 Torben Schinke
 *
 * worldiety Enterprise Edition (EE) License
 * See the file LICENSE for allowed usage and redistribution information.
 *
 * Please contact worldiety GmbH or visit www.worldiety.de if you need additional information or have any
 * questions.
 */

package role

import "github.com/worldiety/suid"

type Role struct {
	ID suid.SUID
}

type Repository interface {
	FindAll(opts struct {
		Limit  int
		Offset int
	}) ([]*Role, error)
	FindById(id suid.SUID) (*Role, error)
	Create() (*Role, error)
	Delete(id suid.SUID) error
}
