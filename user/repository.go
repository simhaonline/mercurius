/*
 * Copyright 2020 Torben Schinke
 *
 * worldiety Enterprise Edition (EE) License
 * See the file LICENSE for allowed usage and redistribution information.
 *
 * Please contact worldiety GmbH or visit www.worldiety.de if you need additional information or have any
 * questions.
 */

package user

import (
	"github.com/worldiety/suid"
)

type Repository interface {

	// Create allocates an empty user, with just an unique Id
	Create() (*User, error)

	// Delete removes the user by id
	Delete(id suid.SUID) error

	// Update takes the given user and overwrites an existing user. It checks if the login is unique (only
	// guaranteed for a single node).
	Update(user *User) error

	// FindAll returns all users in a pageable way
	FindAll(limit int, offset int) ([]*User, error)

	// FindByLogin is a special case, where a user has multiple unique logins
	FindByLogin(login string) (*User, error)
}
