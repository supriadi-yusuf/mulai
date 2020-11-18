package main

import (
	"strings"
	"testing"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func Test_Filter_01(t *testing.T) {

	res, err := simhelper.NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(
		func(i int) bool {
			return i%2 == 1
		})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.([]int)).IsEqual([]int{1, 3, 5})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Filter_02(t *testing.T) {

	res, err := simhelper.NewCollection(
		map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}).FilterValue(
		func(i int) bool {
			return i%2 == 1
		})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.(map[string]int)).IsEqual(
		map[string]int{"one": 1, "three": 3, "five": 5})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Filter_03(t *testing.T) {

	res, err := simhelper.NewCollection([]string{"t1", "t2", "3", "t4", "5", "6"}).FilterValue(
		func(i string) bool {
			return strings.Contains(i, "t")
		})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.([]string)).IsEqual([]string{"t1", "t2", "t4"})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Filter_04(t *testing.T) {

	res, err := simhelper.NewCollection(
		map[string]string{"one": "t1", "two": "t2", "three": "3", "four": "t4", "five": "5"}).FilterValue(
		func(i string) bool {
			return strings.Contains(i, "t")
		})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	resbol, err := simhelper.NewCollection(res.(map[string]string)).IsEqual(
		map[string]string{"one": "t1", "two": "t2", "four": "t4"})
	if err != nil {
		t.Errorf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Filter_05(t *testing.T) {

	_, err := simhelper.NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(i float32) bool {
		return i > 1
	})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}

}

func Test_Filter_06(t *testing.T) {

	_, err := simhelper.NewCollection(
		map[string]string{"one": "t1", "two": "t2", "three": "3", "four": "t4", "five": "5"}).FilterValue(
		func(i float32) bool {
			return i > 1
		})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}

}
