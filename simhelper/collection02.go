package simhelper

import (
	"reflect"
)

func (c *myCollection) filterValueInMap(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newMap := reflect.MakeMap(reflect.TypeOf(c.data))

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		outPrms := funcValue.Call([]reflect.Value{iter.Value()})
		if !outPrms[0].Bool() {
			continue
		}

		newMap.SetMapIndex(iter.Key(), iter.Value())
	}

	return newMap.Interface(), nil
}

func (c *myCollection) filterValueInSliceOfArray(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newSlice := reflect.New(reflect.TypeOf(c.data)).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		outPrms := funcValue.Call([]reflect.Value{dataValue.Index(i)})
		if !outPrms[0].Bool() {
			continue
		}

		newSlice = reflect.Append(newSlice, dataValue.Index(i))
	}

	return newSlice.Interface(), nil
}
