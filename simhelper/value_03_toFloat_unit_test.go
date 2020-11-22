package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_value_03(t *testing.T) {

	res, err := simhelper.NewValue(int(10)).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != 10 {
		t.Errorf("wrong result\n")
	}

	res, err = simhelper.NewValue(int(-10)).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != -10 {
		t.Errorf("wrong result\n")
	}

	res, err = simhelper.NewValue(1.8).ToFloat()
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res != 1.8 {
		t.Errorf("wrong result\n")
	}

	_, err = simhelper.NewValue("test").ToFloat()
	if err == nil {
		t.Errorf("it should be error\n")
	}

}
