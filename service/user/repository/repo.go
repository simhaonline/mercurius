package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/worldiety/mercurius/service/user/repository/mysql"
	"github.com/worldiety/mercurius/service/user/repository/postgresql"
)

type CreateUserParams = struct {
	ID        uuid.UUID      `json:"id"`
	Firstname string         `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Login     string         `json:"login"`
}

type User = struct {
	ID        uuid.UUID      `json:"id"`
	Firstname string         `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Login     string         `json:"login"`
}

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ListUsers(ctx context.Context) ([]User, error)
}

func WrapMysql(q mysql.Querier) Querier {
	return mysqlAdapter{q}
}


type mysqlAdapter struct {
	mysql.Querier
}

func (m mysqlAdapter) CreateUser(ctx context.Context, arg CreateUserParams) error {
	return m.Querier.CreateUser(ctx, arg)
}

func (m mysqlAdapter) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return m.Querier.DeleteUser(ctx, id)
}

func (m mysqlAdapter) ListUsers(ctx context.Context) ([]User, error) {
	res, err := m.Querier.ListUsers(ctx)
	tmp := make([]User, cap(res), len(res))
	for i, a := range res {
		tmp[i] = a
	}
	return tmp, err
}

func WrapPostgresql(q postgresql.Querier) Querier {
	return postgresqlAdapter{q}
}

type postgresqlAdapter struct {
	postgresql.Querier
}

func (m postgresqlAdapter) CreateUser(ctx context.Context, arg CreateUserParams) error {
	return m.Querier.CreateUser(ctx, arg)
}

func (m postgresqlAdapter) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return m.Querier.DeleteUser(ctx, id)
}

func (m postgresqlAdapter) ListUsers(ctx context.Context) ([]User, error) {
	res, err := m.Querier.ListUsers(ctx)
	tmp := make([]User, cap(res), len(res))
	for i, a := range res {
		tmp[i] = a
	}
	return tmp, err
}
