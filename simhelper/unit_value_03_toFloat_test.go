package simhelper

import (
	"log"
	"testing"
)

func Test_ToFloat_01(t *testing.T) {

	log.Println(t.Name())

	res, err := NewValue(int(10)).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != float64(10) {
		t.Errorf("wrong result\n")
	}

	res, err = NewValue(int(-10)).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != float64(-10) {
		t.Errorf("wrong result\n")
	}

	res, err = NewValue(1.8).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != float64(1.8) {
		t.Errorf("wrong result\n")
	}

	_, err = NewValue("test").ToFloat()
	if err == nil {
		t.Errorf("it should be error\n")
	}

}
