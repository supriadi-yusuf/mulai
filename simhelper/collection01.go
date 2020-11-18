package simhelper

import (
	"reflect"
)

func (c *myCollection) removeIndexInMap(index interface{}) (result interface{}, err error) {

	newMap := reflect.MakeMap(reflect.TypeOf(c.data))

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {
		key := iter.Key()
		if key.Type() != reflect.TypeOf(index) {
			panic("map key and index have different data type")
		}

		if key.Interface() == index {
			continue
		}

		newMap.SetMapIndex(key, iter.Value())
	}

	return newMap.Interface(), nil
}

func (c *myCollection) removeIndexInSliceOfArray(index int) (result interface{}, err error) {

	newSlice := reflect.New(reflect.TypeOf(c.data)).Elem()

	dataValue := reflect.ValueOf(c.data)
	for i := 0; i < dataValue.Len(); i++ {
		if i == index {
			continue
		}

		newSlice = reflect.Append(newSlice, dataValue.Index(i))
	}

	return newSlice.Interface(), nil
}
