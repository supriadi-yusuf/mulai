package simhelper_test

import (
	"fmt"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func ExampleGetErrorOnPanic() {

	result := func() (err error) {

		defer simhelper.GetErrorOnPanic(&err)

		fmt.Println("Hello")

		return
	}()
	if result != nil {
		fmt.Println(result.Error())
	}
}
