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

type structtest struct {
	Name   string // beginning of field name must be capital
	Age    int
	Height float32
}

func Test_Equal_08(t *testing.T) {

	ps1 := structtest{"iwan", 10, 1.50}
	ps2 := structtest{"iwan", 10, 1.50}
	result, err := simhelper.NewCollection(ps1).IsEqual(ps2)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !result {

		t.Errorf("Wrong result\n")

	}

	ps3 := structtest{"iwan", 1, 1.5}
	result, err = simhelper.NewCollection(ps1).IsEqual(ps3)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if result {

		t.Errorf("Wrong result\n")

	}
}
