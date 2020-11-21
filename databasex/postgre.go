package databasex

import (
	"database/sql"
	"fmt"
	"strings"
)

type postgresDb struct {
	realDb
}

// CreateQueryString is method implementing interface
func (workDb *postgresDb) CreateValuesMark(fieldNum int) (string, error) {

	var arrPrms []string

	for i := 1; i <= fieldNum; i++ {
		arrPrms = append(arrPrms, fmt.Sprintf("$%d", i))
		//arrPrms = append(arrPrms, "?")
	}

	return strings.Join(arrPrms, ","), nil
}

// CreateConnection is method
func (workDb *postgresDb) createConnection(username, password, host, port, dbname, other string,
	maxConnections, maxIdleConnection int) (*sql.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s", username, password, host, port, dbname, other)

	db, err := sql.Open("postgres", connString)
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

// NewPostgre is a function to connect with postgresql database.
// This function returns 2 objects : the first one has type of IDatabase (interface), the second one has type of error.
//
// This function has several input parameters :
//
// - username is username of database we want to access
//
// - password is password of username
//
// - host is location where postgresql lives
//
// - port is database port
//
// - dbname is name of database
//
// - other is additional parameter if we need it. for example : sslmode=disable
func NewPostgre(username, password, host, port, dbname, other string,
	maxConnections, maxIdleConnection int) (IDatabase, error) {

	var workDb postgresDb

	_, err := workDb.createConnection(username, password, host, port, dbname, other,
		maxConnections, maxIdleConnection)
	if err != nil {
		return nil, err
	}

	return &workDb, nil
}
