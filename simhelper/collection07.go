package simhelper

import (
	"reflect"
)

func (c *myCollection) isElemEqualInMap(data interface{}) (result bool, err error) {

	return c.isEqualInMap(data)
}

func (c *myCollection) findIn(value reflect.Value, collection reflect.Value) (res bool) {

	interfaceVal := value.Interface()
	for i := 0; i < collection.Len(); i++ {
		if interfaceVal == collection.Index(i).Interface() {
			return true
		}

	}

	return false
}

func (c *myCollection) isElemEqualInSliceOfArray(data interface{}) (result bool, err error) {

	paramValue := reflect.ValueOf(data)
	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		if !c.findIn(paramValue.Index(i), dataValue) {
			return false, nil
		}

		if !c.findIn(dataValue.Index(i), paramValue) {
			return false, nil
		}
	}

	return true, nil
}

func (c *myCollection) isElemEqualInStruct(data interface{}) (result bool, err error) {

	return c.isEqualInStruct(data)
}
