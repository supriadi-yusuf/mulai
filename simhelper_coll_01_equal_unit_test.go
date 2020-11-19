package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_Equal_01(t *testing.T) {
	res, err := simhelper.NewCollection([]int{1, 2, 3, 4}).IsEqual([]int{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}
}

func Test_Equal_02(t *testing.T) {
	res, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsEqual(
		map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}
}

func Test_Equal_03(t *testing.T) {

	res, err := simhelper.NewCollection([]int{1, 2, 3, 4}).IsEqual([]int{1, 2, 4, 3})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res {
		t.Errorf("data are not equal")
	}

}

func Test_Equal_04(t *testing.T) {
	res, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsEqual(
		map[string]int{"satu": 1, "dua": 2, "tiga": 4, "empat": 3})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res {
		t.Errorf("data are not equal")
	}
}

func Test_Equal_05(t *testing.T) {

	res, err := simhelper.NewCollection([]string{"father", "mother", "son"}).IsEqual(
		[]string{"father", "mother", "son"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}
}

func Test_Equal_06(t *testing.T) {
	_, err := simhelper.NewCollection([]int{1, 2, 3, 4}).IsEqual([]float32{1, 2, 3, 4})
	if err == nil {
		t.Errorf("type are different")
	}
}

func Test_Equal_07(t *testing.T) {
	_, err := simhelper.NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsEqual(
		map[string]float32{"satu": 1, "dua": 2, "tiga": 3, "empat": 4})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}
}
