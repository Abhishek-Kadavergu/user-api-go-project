package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID   int32
	Name string
	Dob  pgtype.Date
}
