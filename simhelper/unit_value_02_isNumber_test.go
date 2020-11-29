package simhelper

import (
	"log"
	"testing"
)

func Test_IsNumber_01(t *testing.T) {

	log.Println(t.Name())

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
