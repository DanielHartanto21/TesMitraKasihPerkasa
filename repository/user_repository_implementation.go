package repository

import (
	"context"
	"database/sql"
	"errors"
	"tesMitraKasihPerkasa/helper"
	"tesMitraKasihPerkasa/model/domain"
)

type UserRepositoryImplementation struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplementation{}
}
func (users *UserRepositoryImplementation) Logins(ctx context.Context, tx *sql.Tx, user domain.Users) (domain.Users, error) {

	SQL := "select Uuid from Users where NamaUser=$1 and Password=$2"
	rows, err := tx.QueryContext(ctx, SQL, user.NamaUser, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()
	data := domain.Users{NamaUser: user.NamaUser}
	if rows.Next() {
		err := rows.Scan(&data.Uuid)
		helper.PanicIfError(err)
		return data, nil
	} else {
		return data, errors.New("not found")
	}
}
func (users *UserRepositoryImplementation) Check(tx *sql.Tx, uuid string) bool {

	SQL := "select * from Users where Uuid=$1"
	rows, err := tx.Query(SQL, uuid)
	helper.PanicIfError(err)
	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
}
