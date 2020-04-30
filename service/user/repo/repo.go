package repo

import (
	"context"
	"database/sql"
)

//TODO private fields only work within the package itself

type User struct {
	id    string
	login string
	mails []string
}

func (u User) Mails() []string {
	return u.mails
}

func (u User) Login() string {
	return u.login
}

func (u User) Id() string {
	return u.id
}

//TODO placeholder syntax is database-specific, e.g. ? for MYSQL but PostgreSQL is $1, $2 etc.

// v0: CREATE TABLE user (id UUID, login varchar NOT NULL) PRIMARY KEY(id)
// v1: ALTER TABLE ADD name VARCHAR
type SQLRepository interface {
	// SELECT 1 FROM $table WHERE login = :login
	HasLogin(login string) (bool, error)

	// INSERT INTO $table (login) VALUES (:login)
	CreateUser(login string) (User, error)

	// SELECT * FROM $table ORDER BY login DESC LIMIT :limit OFFSET :offset
	All(offset, limit int) ([]User, error)

	// SELECT blub WHERE asdf = :blub
	Blub(ctx context.Context, blub string) error // TODO pollute all the things with context? what about TX and DB dependency?
}

// TODO for request based things, codegen creates a constructor? Or should we pollute the methods with context?
func NewSQLRepository(ctx context.Context, tx *sql.Tx) SQLRepository {
	return nil
}
