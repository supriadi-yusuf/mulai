package databasex_test

import (
	"context"
	"log"

	"github.com/supriadi-yusuf/mulai/databasex"
)

func ExampleNewSimpleSQL_insertConn() {

	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 10, 5)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// get new connection
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	sqlOp := databasex.NewSimpleSQL(postgres)

	// prepare data
	student := Student{"C001", "junjun", 6, 1}

	// create model
	model := databasex.NewSimpleModel("tb_student", student)

	// insert data into table
	if err = sqlOp.InsertConn(context.Background(), conn, model); err != nil {
		log.Fatalln(err.Error())
	}

}

func ExampleNewSimpleSQL_deleteConn() {

	// connect to postgresql with database db_belajar_golang
	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// get database connection
	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// get new connection
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// create sql operatio object
	sqlOp := databasex.NewSimpleSQL(postgres)

	// create model we want to work with
	model := databasex.NewSimpleModel("tb_student", nil)

	// delete record for student whose name is agus
	if _, err = sqlOp.DeleteConn(context.Background(), conn, model, "name='agus'"); err != nil {
		log.Fatalln(err.Error())
	}

	// delete all records
	if _, err = sqlOp.DeleteConn(context.Background(), conn, model, ""); err != nil {
		log.Fatalln(err.Error())
	}

}

func ExampleNewSimpleSQL_updateConn() {

	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// get new connection
	conn, err := db.Conn(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	sqlOp := databasex.NewSimpleSQL(postgres)

	const (
		cName  = "eko win"
		cAge   = 8
		cGrade = 3
	)

	var keypair = struct {
		Name  string
		Age   int
		Grade int
	}{cName, cAge, cGrade}

	// create model
	model := databasex.NewSimpleModel("tb_student", keypair)

	if _, err = sqlOp.UpdateConn(context.Background(), conn, model, "id='C001'"); err != nil {
		log.Fatalln(err.Error())
	}

}
