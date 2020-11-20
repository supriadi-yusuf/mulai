package simhelper

import (
	"reflect"
)

func (c *myCollection) isEqualInMap(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		if iter.Value().Interface() != paramValue.MapIndex(iter.Key()).Interface() {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isEqualInSliceOfArray(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		if paramValue.Index(i).Interface() != dataValue.Index(i).Interface() {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isEqualInStruct(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.NumField(); i++ {

		if paramValue.Field(i).Interface() != dataValue.Field(i).Interface() {
			return false, nil
		}
	}

	return true, nil
}
