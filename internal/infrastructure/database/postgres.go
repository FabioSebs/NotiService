package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgresDB(host, user, password, dbname string, port int) *sqlx.DB {
	// lol
	connString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return db
}
