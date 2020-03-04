package database

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	DB *sqlx.DB
)

func Open(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
}
