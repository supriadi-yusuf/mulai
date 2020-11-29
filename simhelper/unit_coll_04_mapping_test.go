package simhelper

import (
	"log"
	"strings"
	"testing"
)

func Test_Mapping_01(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(i int) bool {
		return i%2 == 1
	})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.([]bool)).IsEqual([]bool{true, false, true, false, true, false})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Mapping_02(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection(
		map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}).MappingValue(
		func(i int) bool {
			return i%2 == 1
		})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.(map[string]bool)).IsEqual(
		map[string]bool{"one": true, "two": false, "three": true, "four": false})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Mapping_03(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection([]string{"s1", "s2", "3", "s4"}).MappingValue(
		func(i string) string {
			if strings.HasPrefix(i, "s") {
				return i
			}

			return "0" + i
		})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.([]string)).IsEqual([]string{"s1", "s2", "03", "s4"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Mapping_04(t *testing.T) {

	log.Println(t.Name())

	res, err := NewCollection(map[int]string{1: "s1", 2: "s2", 3: "3", 4: "s4"}).MappingValue(
		func(i string) string {
			if strings.HasPrefix(i, "s") {
				return i
			}

			return "0" + i
		})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	resbol, err := NewCollection(res.(map[int]string)).IsEqual(
		map[int]string{1: "s1", 2: "s2", 3: "03", 4: "s4"})
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}

	if !resbol {
		t.Errorf("%s\n", "wrong result")
	}

}

func Test_Mapping_05(t *testing.T) {

	log.Println(t.Name())

	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(i float32) bool {
		return i > 1
	})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}

}

func Test_Mapping_06(t *testing.T) {

	log.Println(t.Name())

	_, err := NewCollection(
		map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}).MappingValue(
		func(i float32) bool {
			return i > 1
		})
	if err == nil {
		t.Errorf("%s\n", err.Error())
	}

}
