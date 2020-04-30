// Code generated by sqlc. DO NOT EDIT.

package postgresql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Firstname string         `json:"firstname"`
	Lastname  sql.NullString `json:"lastname"`
	Login     string         `json:"login"`
	CreatedAt time.Time      `json:"created_at"`
}
