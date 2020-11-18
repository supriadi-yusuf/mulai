package simhelper

import (
	"reflect"
)

func (c *myCollection) mappingValueInMap(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newMapType := reflect.MapOf(reflect.TypeOf(c.data).Key(), funcValue.Type().Out(0))
	newMap := reflect.MakeMap(newMapType)

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {

		outPrms := funcValue.Call([]reflect.Value{iter.Value()})

		newMap.SetMapIndex(iter.Key(), outPrms[0])
	}

	return newMap.Interface(), nil
}

func (c *myCollection) mappingValueInSliceOfArray(fcriteria interface{}) (result interface{}, err error) {

	funcValue := reflect.ValueOf(fcriteria)

	newSliceType := reflect.SliceOf(funcValue.Type().Out(0))
	newSlice := reflect.New(newSliceType).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {

		outPrms := funcValue.Call([]reflect.Value{dataValue.Index(i)})

		newSlice = reflect.Append(newSlice, outPrms[0])
	}

	return newSlice.Interface(), nil
}
