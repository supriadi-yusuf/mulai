package databasex_test

import (
	"log"

	"github.com/supriadi-yusuf/mulai/databasex"
)

func ExampleNewPostgre() {
	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

}
