package simhelper

import (
	"testing"
)

func Test_value_02(t *testing.T) {

	if !NewValue(10).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if !NewValue(100.8).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if !NewValue(-10).IsNumber() {
		t.Errorf("wrong result\n")
	}

	if NewValue("hello").IsNumber() {
		t.Errorf("wrong result\n")
	}

}
