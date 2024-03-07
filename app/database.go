package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"tesMitraKasihPerkasa/helper"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "TesMKP"
)

func NewDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	return db
}
