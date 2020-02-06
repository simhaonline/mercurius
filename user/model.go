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

// A User represents a known and managed user
type User struct {
	ID     suid.SUID // ID is unique for the entity
	Logins []string  // Logins contains globally unique names across all users. A user may have multiple logins
	PHash  []byte    // PHash contains a salted and hashed password. We never persist any password.

}

type UsersResponse struct{
	Users []*User
}