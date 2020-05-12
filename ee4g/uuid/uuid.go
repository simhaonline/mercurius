package uuid

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (u *UUID) Scan(src interface{}) error {
	t := (*uuid.UUID)(u)
	return t.Scan(src)
}

func (u UUID) Value() (driver.Value, error) {
	t := [16]byte(u)
	return t[:], nil
}

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

func New() UUID {
	return UUID(uuid.New())
}

func MustParse(str string) UUID {
	return UUID(uuid.MustParse(str))
}
