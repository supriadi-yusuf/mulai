package databasex_test

import (
	"context"
	"fmt"
	"log"

	"github.com/supriadi-yusuf/mulai/databasex"
)

func Example_insert() {

	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 100, 5)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlOp := databasex.NewSimpleSQL(postgres)

	// prepare data
	student := Student{"C001", "junjun", 6, 1}

	// create model
	model := databasex.NewSimpleModel("tb_student", student)

	// insert data into table
	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		log.Fatalln(err.Error())
	}

}

func Example_delete() {

	// connect to postgresql with database db_belajar_golang
	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 100, 5)
	if err != nil {
		log.Fatal(err)
	}

	// get database connection
	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create sql operatio object
	sqlOp := databasex.NewSimpleSQL(postgres)

	// create model we want to work with
	model := databasex.NewSimpleModel("tb_student", nil)

	// delete record for student whose name is agus
	if _, err = sqlOp.DeleteDb(context.Background(), model, "name='agus'"); err != nil {
		log.Fatalln(err.Error())
	}

	// delete all records
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		log.Fatalln(err.Error())
	}

}

func Example_update() {

	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 100, 5)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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

	if _, err = sqlOp.UpdateDb(context.Background(), model, "id='C001'"); err != nil {
		log.Fatalln(err.Error())
	}

}

func Example_select() {

	type Student struct {
		ID    string
		Name  string
		Age   int
		Grade int
	}

	postgres, err := databasex.NewPostgre("scott", "tiger", "localhost", "5432", "db_belajar_golang",
		"?sslmode=disable", 100, 5)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlOp := databasex.NewSimpleSQL(postgres)

	// create model
	model := databasex.NewSimpleModel("tb_student", Student{})

	data := make([]Student, 0)

	// select record for ID = C001
	if err = sqlOp.SelectDb(context.Background(), model, "ID='C001'", &data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(data)

	data = make([]Student, 0)

	// select all records
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(data)

}
