package repository

import (
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImplementation struct {
}

func (users *UserRepositoryImplementation) Login(ctx context.Context, tx *sql.Tx, name string, password string) (string, error) {
	//TODO implement me
	panic("implement me")
	SQL := "select Uuid from Users where NamaUser=?,Password=?"
	rows, err := tx.ExecContext(ctx, SQL, name, password)
	if err != nil {
		return "", err
	}
	if rows.Next() {
		rows.Scan(Uuid)
	} else {
		return "", errors.New("not found")
	}
}
