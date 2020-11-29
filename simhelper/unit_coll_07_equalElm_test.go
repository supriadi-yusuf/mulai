package simhelper

import (
	"log"
	"testing"
)

func Test_EqualElem_01(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]int{1, 2, 3, 4}).IsElemEqual([]int{1, 2, 3, 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}
}

func Test_EqualElem_02(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsElemEqual(
		map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}
}

func Test_EqualElem_03(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]int{1, 2, 3, 4}).IsElemEqual([]int{1, 2, 4, 3})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}

}

func Test_EqualElem_04(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsElemEqual(
		map[string]int{"satu": 1, "dua": 2, "tiga": 4, "empat": 3})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res {
		t.Errorf("data are not equal")
	}
}

func Test_EqualElem_05(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]string{"father", "mother", "son"}).IsElemEqual(
		[]string{"father", "mother", "son"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}

	res, err = NewCollection([]string{"father", "mother", "son"}).IsElemEqual(
		[]string{"mother", "father", "son"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("data are equal")
	}

	res, err = NewCollection([]string{"father", "mother"}).IsElemEqual(
		[]string{"mother", "father", "son"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if res {
		t.Errorf("data are not equal")
	}
}

func Test_EqualElem_06(t *testing.T) {

	log.Println(t.Name())

	_, err := NewCollection([]int{1, 2, 3, 4}).IsElemEqual([]float32{1, 2, 3, 4})
	if err == nil {
		t.Errorf("type are different")
	}
}

func Test_EqualElem_07(t *testing.T) {

	log.Println(t.Name())

	_, err := NewCollection(map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).IsElemEqual(
		map[string]float32{"satu": 1, "dua": 2, "tiga": 3, "empat": 4})
	if err == nil {
		t.Errorf("type are different")
	}
}

func Test_EqualElem_08(t *testing.T) {

	log.Println(t.Name())

	type structtest struct {
		Name   string // beginning of field name must be capital
		Age    int
		Height float32
	}

	ps1 := structtest{"iwan", 10, 1.50}
	ps2 := structtest{"iwan", 10, 1.50}
	result, err := NewCollection(ps1).IsElemEqual(ps2)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !result {

		t.Errorf("Wrong result\n")

	}

	ps3 := structtest{"iwan", 1, 1.5}
	result, err = NewCollection(ps1).IsElemEqual(ps3)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if result {

		t.Errorf("Wrong result\n")

	}
}
