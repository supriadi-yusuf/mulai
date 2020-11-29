package simhelper

import (
	"fmt"
)

//GetErrorOnPanic is function. This function checks weither panic is happen or not.
//if panic happen, the panic will be caught and become an error
func GetErrorOnPanic(err *error) {
	//if *err != nil {
	//	return
	//}

	if it := recover(); it != nil {
		*err = fmt.Errorf("%v", it)
	}
}
