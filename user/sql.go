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
	"database/sql"
	"github.com/worldiety/mercurius/ee"
	"github.com/worldiety/suid"
)

var _ Repository = (*sqlRepository)(nil)

type sqlRepository struct {
}

func NewSQLRepository(tx *sql.Tx) Repository {
	return nil
}

func (r *sqlRepository) Create() (*User, error) {
	panic("implement me")
}

func (r *sqlRepository) Delete(id suid.SUID) error {
	panic("implement me")
}

func (r *sqlRepository) Update(user *User) error {
	panic("implement me")
}

func (r *sqlRepository) FindAll(ctx *ee.Ctx, opts struct {
	Limit  int
	Offset int
}) ([]*User, error) {
	var users []*User

	err := ctx.
		SQL().Query("SELECT * FROM user LIMIT ? OFFSET ?", opts.Limit, opts.Offset).
		Map(func(row ee.Row) error {
			user := &User{}
			users = append(users, user)
			return row.Scan(&user.ID, &user.PHash)
		})

	return users, err
}

func (r *sqlRepository) FindByLogin(ctx *ee.Ctx, login string) (*User, error) {
	panic("implement me")
}
