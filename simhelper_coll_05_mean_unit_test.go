package main

import (
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_Mean_01(t *testing.T) {
	res, err := simhelper.NewCollection([]int{1, 2, 3, 4}).MeanValue()
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res != 2.5 {
		t.Errorf("wrong result")
	}
}

func Test_Mean_02(t *testing.T) {
	res, err := simhelper.NewCollection(
		map[string]int{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).MeanValue()
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res != 2.5 {
		t.Errorf("wrong result")
	}
}

func Test_Mean_03(t *testing.T) {
	res, err := simhelper.NewCollection([]float32{1, 2, 3, 4}).MeanValue()
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res != 2.5 {
		t.Errorf("wrong result")
	}
}

func Test_Mean_04(t *testing.T) {
	res, err := simhelper.NewCollection(
		map[string]float32{"satu": 1, "dua": 2, "tiga": 3, "empat": 4}).MeanValue()
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if res != 2.5 {
		t.Errorf("wrong result")
	}
}

func Test_Mean_05(t *testing.T) {
	_, err := simhelper.NewCollection([]string{"a", "b", "c"}).MeanValue()
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}
}

func Test_Mean_06(t *testing.T) {
	_, err := simhelper.NewCollection(map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}).MeanValue()
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}
}
