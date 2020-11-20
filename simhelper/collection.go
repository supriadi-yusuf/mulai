package simhelper

import (
	"reflect"
)

type myCollection struct {
	data interface{}
}

func (c *myCollection) RemoveIndex(index interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	if collectionType.Kind() == reflect.Map {
		return c.removeIndexInMap(index)
	}

	return c.removeIndexInSliceOfArray(index.(int))
}

func (c *myCollection) FilterValue(fcriteria interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	//fcriteria must function with 1 input argument and 1 output argument.
	//output argument must be boolean
	ftype := reflect.TypeOf(fcriteria)
	if ftype.Kind() != reflect.Func {
		panic("fcriteria argument must be function")
	}

	if ftype.NumIn() != 1 {
		panic("fcritera must have only one input parameter")
	}

	if ftype.In(0) != collectionType.Elem() {
		panic("element of collection and input parameter for fcriteria are different int type")
	}

	if ftype.NumOut() != 1 {
		panic("fcritera must have only one output parameter")
	}

	if ftype.Out(0) != reflect.TypeOf(true) {
		panic("output parameter for fcriteria must be boolean")
	}

	if collectionType.Kind() == reflect.Map {
		return c.filterValueInMap(fcriteria)
	}

	return c.filterValueInSliceOfArray(fcriteria)
}

//Mapping is map every value on collection based on function fmapping.
func (c *myCollection) MappingValue(fmapping interface{}) (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	//fcriteria must function with 1 input argument and 1 output argument.
	ftype := reflect.TypeOf(fmapping)
	if ftype.Kind() != reflect.Func {
		panic("fcriteria argument must be function")
	}

	if ftype.NumIn() != 1 {
		panic("fcritera must have only one input parameter")
	}

	if ftype.In(0) != collectionType.Elem() {
		panic("element of collection and input parameter for fcriteria are different int type")
	}

	if ftype.NumOut() != 1 {
		panic("fcritera must have only one output parameter")
	}

	if collectionType.Kind() == reflect.Map {
		return c.mappingValueInMap(fmapping)
	}

	return c.mappingValueInSliceOfArray(fmapping)
}

//Mean is count average of data in collection
func (c *myCollection) MeanValue() (result float64, err error) {

	defer GetErrorOnPanic(&err)

	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	if collectionType.Kind() == reflect.Map {
		return c.meanValueFromMap()
	}

	return c.meanValueFromSliceOfArray()
}

func (c *myCollection) IsEqual(data interface{}) (result bool, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	//check type of input parameter
	paramType := reflect.TypeOf(data)
	if paramType.Kind() != reflect.Array && paramType.Kind() != reflect.Slice &&
		paramType.Kind() != reflect.Map && paramType.Kind() != reflect.Struct {
		panic("input param must be array, slice, struct or map")
	}

	if paramType != collectionType {
		panic("types are different")
	}

	collectionValue := reflect.ValueOf(c.data)
	paramValue := reflect.ValueOf(data)
	if collectionValue.Len() != paramValue.Len() {
		return false, nil
	}

	if collectionType.Kind() == reflect.Map {
		return c.isEqualInMap(data)
	}

	if collectionType.Kind() == reflect.Struct {
		return c.isEqualInStruct(data)
	}

	return c.isEqualInSliceOfArray(data)

}

func (c *myCollection) ConvElmToInterface() (result interface{}, err error) {

	defer GetErrorOnPanic(&err)

	//check type of collection's data
	collectionType := reflect.TypeOf(c.data)
	if collectionType.Kind() != reflect.Array && collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		panic("collection must be array, slice or map")
	}

	if collectionType.Kind() == reflect.Map {
		return c.convElmToInterfaceOnMap()
	}

	return c.convElmToInterfaceOnSliceOrArray()
}

//NewCollection is function creating object whose type is ICollection
func NewCollection(data interface{}) ICollection {
	collection := new(myCollection)
	collection.data = data
	return collection
}
