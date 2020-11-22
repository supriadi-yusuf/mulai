package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_Convert_01(t *testing.T) {

	res, err := simhelper.NewCollection([]int{1, 2, 3, 4, 5, 6}).ConvElmToInterface()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.([]interface{})).IsEqual([]interface{}{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Convert_02(t *testing.T) {

	res, err := simhelper.NewCollection(
		map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}).ConvElmToInterface()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.(map[string]interface{})).IsEqual(
		map[string]interface{}{"one": 1, "two": 2, "three": 3, "four": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}
