package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/supriadi-yusuf/mulai/databasex"
)

/*
const (
	mysqlUsername       = "root"
	mysqlPassword       = ""
	mysqlHost           = "localhost"
	mysqlPort           = "3306"
	mysqlDb             = "db_belajar_golang"
	mysqlMaxConnections = 0
	mysqlMaxIdle        = 0
)
*/
// Student is type
/*type Student struct {
	ID    string
	Name  string
	Age   int
	Grade int
}*/

func TestCreateTableMySqlCommit(t *testing.T) {

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	cmdStr := "drop table if exists tb_student"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	cmdStr = "create table if not exists tb_student(id varchar(5),	name varchar(255),age int,grade int)"
	_, err = db.Exec(cmdStr)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

}

func TestAddOneRecordMySqlCommit(t *testing.T) {

	t.Logf("testing : add one record to tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertTrans(context.Background(), tx, model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, fmt.Sprintf("ID='%s'", student.ID), &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) < 1 {
		t.Errorf("adding one data fail")
	}

	if data[0].ID != student.ID || data[0].Name != student.Name || data[0].Age != student.Age || data[0].Grade != student.Grade {
		t.Errorf("data is different")
	}

}

func TestUpdateOneRecordMySqlCommit(t *testing.T) {

	t.Logf("testing : update one records from tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName  = "eko win"
		CAge   = 8
		CGrade = 3
	)

	var keypair = struct {
		Name  string
		Age   int
		Grade int
	}{CName, CAge, CGrade}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateTrans(context.Background(), tx, model, "id='C001'"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "id='C001'", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if data[0].Name != CName || data[0].Age != CAge || data[0].Grade != CGrade {
		t.Errorf("data is different")
	}

}

func TestDeleteOneRecordMySqlCommit(t *testing.T) {

	t.Logf("testing : delete one records from tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("insert one record into table")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("delete one records from table")

	model.SetNewData(nil)
	if _, err = sqlOp.DeleteTrans(context.Background(), tx, model, "id='C001'"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read from table")

	data := make([]Student, 0)

	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "id='C001'", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) != 0 {
		t.Errorf("data is not deleted")
	}

}

func TestUpdateSeveralRecordsMySqlCommit(t *testing.T) {

	t.Logf("testing : update several records from tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName = "eko win"
		CAge  = 8
	)

	var keypair = struct {
		Name string
		Age  int
	}{CName, CAge}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateTrans(context.Background(), tx, model, "grade=5"); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "grade=5", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) == 0 {
		t.Errorf("data is not inserted")
	}

	for _, val := range data {

		if val.Name != CName || val.Age != CAge {
			t.Errorf("data is different")
			break
		}

	}

}

func TestUpdateAllRecordsMySqlCommit(t *testing.T) {

	t.Logf("testing : update all records from tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("update record")

	const (
		CName = "eko win"
		CAge  = 8
	)

	var keypair = struct {
		Name string
		Age  int
	}{CName, CAge}

	model.SetNewData(keypair)
	if _, err = sqlOp.UpdateTrans(context.Background(), tx, model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) == 0 {
		t.Errorf("data is not inserted")
	}

	for _, val := range data {

		if val.Name != CName || val.Age != CAge {
			t.Errorf("data is different")
			break
		}

	}

}

func TestDeleteAllRecordsMySqlCommit(t *testing.T) {

	t.Logf("testing : delete all records from tabel tb_student in db_belajar_golang database using mysql")

	t.Logf("create connection to database server")

	currDb, err := databasex.NewMysql(mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDb,
		mysqlMaxConnections, mysqlMaxIdle)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	db, err := currDb.GetDbConnection()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	sqlOp := databasex.NewSimpleSQL(currDb)

	t.Logf("delete all data first")

	model := databasex.NewSimpleModel("tb_student", nil)
	if _, err = sqlOp.DeleteDb(context.Background(), model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("add several data")

	student := Student{"C001", "junjun", 6, 1}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C002", "maman", 8, 2}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	student = Student{"C003", "yuli", 10, 5}
	model.SetNewData(student)

	if err = sqlOp.InsertDb(context.Background(), model); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	t.Logf("delete all records from table")

	model.SetNewData(nil)
	if _, err = sqlOp.DeleteTrans(context.Background(), tx, model, ""); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	tx.Commit()

	t.Logf("read from table")

	data := make([]Student, 0)
	model.SetNewData(Student{})
	if err = sqlOp.SelectDb(context.Background(), model, "", &data); err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if len(data) != 0 {
		t.Errorf("table still has data")
	}

}
