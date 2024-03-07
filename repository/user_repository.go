package repository

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	Login(ctx context.Context, tx *sql.Tx, name string, password string) (string, error)
}
