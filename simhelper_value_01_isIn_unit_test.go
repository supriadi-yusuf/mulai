package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_value_01(t *testing.T) {

	res, err := simhelper.NewValue(int(10)).IsIn([]int{1, 2, 3, 4, 5, 10})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("value is in slice\n")
	}

	res, err = simhelper.NewValue(byte(10)).IsIn([]byte{1, 2, 10, 4, 5, 6})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("value is in slice\n")
	}

	res, err = simhelper.NewValue(float32(10)).IsIn([]float32{1, 2, 3, 4, 5, 6})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res {
		t.Errorf("value is not in slice\n")
	}

	res, err = simhelper.NewValue(int(10)).IsIn([]float32{1, 2, 3, 4, 5, 6})
	if err == nil {
		t.Errorf("type is different\n")
	}

	res, err = simhelper.NewValue("book").IsIn(map[int]string{1: "animal", 2: "house", 3: "giraffe", 4: "book"})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !res {
		t.Errorf("value is in map\n")
	}

	res, err = simhelper.NewValue("book").IsIn(map[int]string{1: "animal", 2: "house", 3: "giraffe", 4: "science"})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res {
		t.Errorf("value is not in map\n")
	}

	res, err = simhelper.NewValue(10).IsIn(map[int]string{1: "animal", 2: "house", 3: "giraffe", 4: "book"})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}

}
