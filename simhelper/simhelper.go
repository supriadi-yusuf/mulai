//Package simhelper is simple helper function to help developer working with data/value.
package simhelper

// IValue is interface defines what to do with single data / value.
//
// This interface has several methods :
//
// - IsIn(collection interface{}) (result bool, err error)
//   Check if certain value is in collection of data.
//   Parameter collection must be array, slice or map
//
// - IsNumber() (result bool)
//
//   Check if certain value is number or not.
//
// - ToFloat() (result float64, err error)
//
//   Converts certain value into float64
//
//
// IValue is declarated as follow :
//
type IValue interface {
	IsIn(collection interface{}) (result bool, err error)
	IsNumber() (result bool)
	ToFloat() (result float64, err error)
}

// ICollection is interface defines what to do with collection of data ( map, slice, array, struct).
//
// This interface has several methods :
//
// - RemoveIndex(index interface{}) (result interface{}, err error)
//   Remove element with certain index from collection (map, array, list).
//
// - FilterValue(fcriteria interface{}) (result interface{}, err error)
//   Filter data in collection (map, array, list) based on input parameter named fcriteria.
//   Parameter fcriteria must be function.
//
// - MappingValue(fmapping interface{}) (result interface{}, err error)
//   Map every value in collection (map, array, list) based on input parameter named fmapping.
//   Parameter fmapping must be function.
//
// - MeanValue() (result float64, err error)
//   Count average of value in collection
//
// - IsEqual(data interface{}) (result bool, err error)
//   Check if two map/slice/struct are equal or not.
//   Parameter data must be map/slice/struct.
//
// - ConvElmToInterface() (result interface{}, err error)
//   Converts each element in slice / map to interface{}
//
// ICollection is declarated as follow :
type ICollection interface {
	RemoveIndex(index interface{}) (result interface{}, err error)

	FilterValue(fcriteria interface{}) (result interface{}, err error)

	MappingValue(fmapping interface{}) (result interface{}, err error)

	MeanValue() (result float64, err error)

	IsEqual(data interface{}) (result bool, err error)

	ConvElmToInterface() (result interface{}, err error)
}
