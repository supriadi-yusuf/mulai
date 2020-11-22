package simhelper

import (
	"errors"
	"reflect"
)

type myValue struct {
	data interface{}
}

func (v *myValue) IsIn(collection interface{}) (result bool, err error) {

	defer GetErrorOnPanic(&err)

	collectionType := reflect.TypeOf(collection)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice && collectionType.Kind() != reflect.Map {
		return false, errors.New("collection must be array, slice or map")
	}

	if collectionType.Elem() != reflect.TypeOf(v.data) {
		return false, errors.New("type is different")
	}

	if collectionType.Kind() == reflect.Map {
		return v.isInMap(collection), nil
	}

	return v.isInSlcOrArray(collection), nil
}

func (v *myValue) IsNumber() (result bool) {

	valType := reflect.TypeOf(v.data)

	switch valType.Name() {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
		return true
	default:
		return false
	}

}

func (v *myValue) ToFloat() (result float64, err error) {

	defer GetErrorOnPanic(&err)

	elmVal := reflect.ValueOf(v.data)

	switch elmVal.Type().Name() {
	case "int", "int8", "int16", "int32", "int64":
		return float64(elmVal.Int()), nil
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return float64(elmVal.Uint()), nil
	default:
		return elmVal.Float(), nil
	}

}

//NewValue is function to create object whose type is IValue.
//
//data parameter must be not slice, array nor map.
func NewValue(data interface{}) IValue {
	var mValue myValue

	mValue.data = data

	return &mValue
}
