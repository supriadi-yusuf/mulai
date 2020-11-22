package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_value_02(t *testing.T) {

	if !simhelper.NewValue(10).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if !simhelper.NewValue(100.8).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if !simhelper.NewValue(-10).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if simhelper.NewValue("hello").IsNumber() {
		t.Errorf("wrong result\n")
	}

}
