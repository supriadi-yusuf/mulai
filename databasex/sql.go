package databasex

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

// ISqlOperation is interface related to sql operation (CRUD). This interface has several methods :
//
// - InsertDb
//
//   This method is to add new data into database table. This has two input parameters.
//   The first one has type of context.Context. The second one has type of IModel.
//
// - DeleteDb
//
//   This method is to delete records from database table. This method has three input parameters.
//   The last one is criteria. If we delete table with criteria then we can use it.
//   If we do not use the criteria just put empty string into the criteria parameter.
//
//   for example, we run delete operation :
//
//   "delete from tb_student where name='charles'""
//
//   name='charles' is criteria that we need to put into criteria parameter.
//
// - UpdateDb
//
//   This method is to update data in database table.
//
// - SelectDb
//
//   This method is to retrieve data from database table.
//   Retrieved data will be stored into result parameter.
//   When this method is called, value that we should store into result parameter must be address of struct slice.
//
// ISqlOperation is defined as :
type ISqlOperation interface {
	// insert data into table
	InsertDb(ctx context.Context, model IModel) error

	// delete data from table
	DeleteDb(ctx context.Context, model IModel, criteria string) (int, error)

	// update data on table
	UpdateDb(ctx context.Context, model IModel, criteria string) (int, error)

	// retrieve data from table
	SelectDb(ctx context.Context, model IModel, criteria string, result interface{}) error
}

type simpleSQL struct {
	db IDatabase
}

func (s *simpleSQL) InsertDb(ctx context.Context, model IModel) error {

	if err := inspectContext(ctx); err != nil {
		return err
	}

	cmdStr, values, err := createInsertCommand( /*ctx,*/ model, s.getDb().createValuesMark)
	if err != nil {
		return err
	}

	db, err := s.getDb().GetDbConnection()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, cmdStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *simpleSQL) DeleteDb(ctx context.Context, model IModel, criteria string) (int, error) {

	if err := inspectContext(ctx); err != nil {
		return 0, err
	}

	cmdStr := fmt.Sprintf("delete from %s", model.GetTableName())
	if criteria != "" {
		cmdStr = fmt.Sprintf("%s where %s", cmdStr, criteria)
	}

	db, err := s.getDb().GetDbConnection()
	if err != nil {
		return 0, err
	}

	rst, err := db.ExecContext(ctx, cmdStr)
	if err != nil {
		return 0, err
	}

	affectRow, _ := rst.RowsAffected()

	return int(affectRow), nil
}

func (s *simpleSQL) UpdateDb(ctx context.Context, model IModel, criteria string) (int, error) {

	if err := inspectContext(ctx); err != nil {
		return 0, err
	}

	cmdStr, err := createUpdateCommand(model)
	if err != nil {
		return 0, err
	}

	//fmt.Println(cmdStr)

	if criteria != "" {
		cmdStr = fmt.Sprintf("%s where %s", cmdStr, criteria)
	}

	db, err := s.getDb().GetDbConnection()
	if err != nil {
		return 0, err
	}

	rst, err := db.ExecContext(ctx, cmdStr)
	if err != nil {
		return 0, err
	}

	affectRow, _ := rst.RowsAffected()

	return int(affectRow), nil
}

func (s *simpleSQL) SelectDb(ctx context.Context, model IModel, criteria string, result interface{}) error {

	if err := inspectContext(ctx); err != nil {
		return err
	}

	rstType, err := inspectResultOfSelect(result)
	if err != nil {
		return err
	}

	cmdStr, err := createSelectCommand(model)
	if err != nil {
		return err
	}

	if criteria != "" {
		cmdStr = fmt.Sprintf("%s where %s", cmdStr, criteria)
	}

	//fmt.Println(cmdStr)

	db, err := s.getDb().GetDbConnection()
	if err != nil {
		return err
	}

	//var rows *sql.Rows

	rows, err := db.QueryContext(ctx, cmdStr)
	if err != nil {
		return err
	}

	defer rows.Close()

	newRst := reflect.New(rstType).Elem()
	//fmt.Println(newRst)

	for rows.Next() {

		structData := reflect.New(rstType.Elem()).Elem()

		inPrms := make([]reflect.Value, 0)

		for i := 0; i < structData.NumField(); i++ {
			inPrms = append(inPrms, structData.Field(i).Addr())
		}

		outPrms := reflect.ValueOf(rows.Scan).Call(inPrms)
		errOut := outPrms[0].Interface()
		if errOut != nil {
			err, ok := errOut.(error)
			if ok {
				return err
			}

			return errors.New("unkown error after scanning")
		}

		newRst = reflect.Append(newRst, structData)
	}

	rstValue := reflect.ValueOf(result) // ptr to slice of struct
	rstValue = rstValue.Elem()          // slice of struct
	rstValue.Set(newRst)

	return nil
}

func (s *simpleSQL) getDb() IDatabase {
	return s.db
}

// NewSimpleSQL is function returning object whose type is ISqlOperation.
// We need the object to do CRUD on database table.
// This function receives one input parameter.
// The parameter is dbms that you are connected with.
func NewSimpleSQL(db IDatabase) ISqlOperation {
	var s simpleSQL

	s.db = db

	return &s
}
