package simhelper

import "reflect"

func (v *myValue) isInMap(mapCollection interface{}) bool {
	iter := reflect.ValueOf(mapCollection).MapRange()
	for iter.Next() {

		if v.data == iter.Value().Interface() {
			return true
		}

	}

	return false
}

func (v *myValue) isInSlcOrArray(slcArr interface{}) bool {

	slcArrValue := reflect.ValueOf(slcArr)
	slcArrTotal := slcArrValue.Len()
	for i := 0; i < slcArrTotal; i++ {

		if v.data == slcArrValue.Index(i).Interface() {
			return true
		}

	}

	return false
}
