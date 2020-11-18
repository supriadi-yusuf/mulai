//Package simhelper is simple helper function to help developer working with data/value
package simhelper

//IValue is interface defines what to do with single data / value
type IValue interface {
	//IsIn is method to check if certain value is in collection of data
	//parameter collection must be array, slice or map
	IsIn(collection interface{}) (result bool, err error)
	IsNumber() (result bool)
	ToFloat() (result float64, err error)
}

//ICollection is interface defines what to do with collection of data ( map, slice, array)
type ICollection interface {

	//Remove element with certain index from collection
	RemoveIndex(index interface{}) (result interface{}, err error)

	//Filter collection by criteria based on fcriteria. fcriteria is function
	FilterValue(fcriteria interface{}) (result interface{}, err error)

	//Mapping is map every value on collection based on function fmapping.
	MappingValue(fmapping interface{}) (result interface{}, err error)

	//Mean is count average of data in collection
	MeanValue() (result float64, err error)

	//IsEqual check if two map/slice are equal or not
	//parameter data must be slice or map
	IsEqual(data interface{}) (result bool, err error)

	//Converts each element in slice / map to interface{}
	ConvElmToInterface() (result interface{}, err error)
}
