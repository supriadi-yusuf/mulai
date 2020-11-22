package simhelper_test

import (
	"fmt"
	"log"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func ExampleNewValue_isIn() {

	found, err := simhelper.NewValue(float32(10)).IsIn([]float32{2, 4, 6, 8, 10, 12})
	if err != nil {
		log.Fatalln(err.Error())
	}

	if found {
		fmt.Println("10 is in []float32{2,4,6,8,10,12}") // Output ==> 10 is in []float32{2,4,6,8,10,12}
	}

}

func ExampleNewValue_isNumber() {

	if simhelper.NewValue(100).IsNumber() {
		fmt.Println("100 is number") // Output ==> 100 is number
	}

}

func ExampleNewValue_toFloat() {

	res, err := simhelper.NewValue(15).ToFloat()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(res)
}
