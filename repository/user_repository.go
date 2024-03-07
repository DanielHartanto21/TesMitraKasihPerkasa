package repository

import (
	"context"
	"database/sql"
	"tesMitraKasihPerkasa/model/domain"
)

type UserRepository interface {
	Logins(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error)
	Check(tx *sql.Tx, uuid string) bool
}
