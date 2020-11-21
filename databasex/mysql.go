package databasex

import (
	"database/sql"
	"fmt"
	"strings"
)

type mysqlDb struct {
	realDb
}

// CreateQueryString is method implementing interface
func (workDb *mysqlDb) CreateValuesMark(fieldNum int) (string, error) {

	var arrPrms []string

	for i := 1; i <= fieldNum; i++ {
		//arrPrms = append(arrPrms, fmt.Sprintf("$%d", i))
		arrPrms = append(arrPrms, "?")
	}

	return strings.Join(arrPrms, ","), nil
}

// CreateConnection is method
func (workDb *mysqlDb) createConnection(username, password, host, port, dbname string,
	maxConnections, maxIdleConnection int) (*sql.DB, error) {
	//"root:@tcp(127.0.0.1:3306)/db_belajar_golang"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if maxConnections != 0 {
		db.SetMaxOpenConns(maxConnections)
		db.SetMaxIdleConns(maxIdleConnection)
	}

	workDb.db = db

	return db, nil
}

// NewMysql is a function to connect with mysql database.
// This function returns 2 objects : the first one has type of IDatabase (interface), the second one has type of error.
//
// This function has several input parameters :
//
// - username is username of database we want to access
//
// - password is password of username
//
// - host is location where mysql lives
//
// - port is database port
//
// - dbname is name of database
//
func NewMysql(username, password, host, port, dbname string,
	maxConnections, maxIdleConnection int) (IDatabase, error) {

	var workDb mysqlDb

	_, err := workDb.createConnection(username, password, host, port, dbname,
		maxConnections, maxIdleConnection)
	if err != nil {
		return nil, err
	}

	return &workDb, nil
}
