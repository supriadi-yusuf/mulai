package simhelper

import (
	"reflect"
)

func (c *myCollection) convElmToInterfaceOnMap() (result interface{}, err error) {

	mapType := reflect.TypeOf(c.data)
	//newMapType := reflect.MapOf(mapType.Key(), reflect.TypeOf(make([]interface{}, 0)).Elem())
	//var inter1 *interface{} = nil
	//newMapType := reflect.MapOf(mapType.Key(), reflect.TypeOf(inter1).Elem())
	newMapType := reflect.MapOf(mapType.Key(), reflect.TypeOf((*interface{})(nil)).Elem())

	newMapVal := reflect.MakeMap(newMapType)

	iter := reflect.ValueOf(c.data).MapRange()
	for iter.Next() {
		newMapVal.SetMapIndex(iter.Key(), reflect.ValueOf(iter.Value().Interface()))
	}

	return newMapVal.Interface(), nil
}

func (c *myCollection) convElmToInterfaceOnSliceOrArray() (result interface{}, err error) {

	newSliceVal := reflect.New(reflect.TypeOf(make([]interface{}, 0))).Elem()

	slc := reflect.ValueOf(c.data)
	for i := 0; i < slc.Len(); i++ {
		newSliceVal = reflect.Append(newSliceVal, reflect.ValueOf(slc.Index(i).Interface()))
	}

	return newSliceVal.Interface(), nil
}
