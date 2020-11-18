package simhelper

import (
	"reflect"
)

func (c *myCollection) meanValueFromMap() (result float64, err error) {

	var total float64
	var currMap = reflect.ValueOf(c.data)

	iter := currMap.MapRange()
	for iter.Next() {

		res, err := NewValue(iter.Value().Interface()).ToFloat()
		if err != nil {
			return 0, err
		}

		total += res
	}

	return total / float64(currMap.Len()), nil
}

func (c *myCollection) meanValueFromSliceOfArray() (result float64, err error) {

	var total float64
	var currSlice = reflect.ValueOf(c.data)

	for i := 0; i < currSlice.Len(); i++ {

		res, err := NewValue(currSlice.Index(i).Interface()).ToFloat()
		if err != nil {
			return 0, err
		}

		total += res
	}

	return total / float64(currSlice.Len()), nil
}
