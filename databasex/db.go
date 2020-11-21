// Package databasex is to help developer for building app with golang especially when need to interact with dbms.
package databasex

import (
	"database/sql"
	"errors"
	//_ "github.com/lib/pq"
)

// IDatabase is interface related to dbms. This interface has several methods :
//
// - CreateValuesMark
//   CreateValuesMark is a method to create values mark in sql statement.
//   If we use sql to insert data into postgresql. we create command :
//   "insert into tb_xxx(field1, field2, field3) values($1,$2,$3)".
//   $1,$2,$3 are values mark on postgresql.
//
//   But if we want to use mysql, our command should be :
//   "insert into tb_xxx(field1, field2, field3) values(?,?,?)"
//   ?,?,? are values mark on mysql.
//
// - GetDbConnection
//   This method is to get dbms's connection.
//
type IDatabase interface {
	CreateValuesMark(fieldNum int) (string, error)
	GetDbConnection() (*sql.DB, error)
}

type realDb struct {
	db *sql.DB
}

// GetConnection is method
func (workDb *realDb) GetDbConnection() (*sql.DB, error) {

	if workDb.db == nil {
		return nil, errors.New("Connection does not exist")
	}

	return workDb.db, nil

}
