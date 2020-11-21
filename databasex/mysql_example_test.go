package databasex_test

import (
	"fmt"
	"log"

	"github.com/supriadi-yusuf/mulai/databasex"
)

func ExampleNewMysql() {
	currDb, err := databasex.NewMysql("root", "", "localhost", "3306", "db_belajar_golang",
		0, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	valuesmark, _ := currDb.CreateValuesMark(5)
	fmt.Println(valuesmark)
	//Output:
	// ?,?,?,?,?

}
