package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_Remove_01(t *testing.T) {
	res, err := simhelper.NewCollection([]int{1, 2, 3, 4}).RemoveIndex(0)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resBool, err := simhelper.NewCollection(res.([]int)).IsEqual([]int{2, 3, 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resBool {
		t.Errorf("wrong result")
	}
}

func Test_Remove_02(t *testing.T) {
	res, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).RemoveIndex("dua")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resBool, err := simhelper.NewCollection(res.(map[string]int)).IsEqual(map[string]int{"satu": 1, "tiga": 3, "empat": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resBool {
		t.Errorf("wrong result")
	}
}

func Test_Remove_03(t *testing.T) {

	res, err := simhelper.NewCollection([]string{"s1", "s2", "s3", "s4"}).RemoveIndex(4)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resBool, err := simhelper.NewCollection(res.([]string)).IsEqual([]string{"s1", "s2", "s3", "s4"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resBool {
		t.Errorf("wrong result")
	}

}

func Test_Remove_04(t *testing.T) {
	res, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).RemoveIndex(
		"lima")
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resBoll, err := simhelper.NewCollection(res.(map[string]int)).IsEqual(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resBoll {
		t.Errorf("wrong result")
	}
}

func Test_Remove_05(t *testing.T) {
	_, err := simhelper.NewCollection([]int{1, 2, 3, 4}).RemoveIndex(10.5)
	if err == nil {
		t.Errorf("index has wrong type")
	}
}

func Test_Remove_06(t *testing.T) {
	_, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).RemoveIndex(0)
	if err == nil {
		t.Errorf("index has different type")
	}
}
