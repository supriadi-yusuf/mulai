package simhelper

import (
	"testing"
)

func sendValInt(unity []int) (ch <-chan int) {

	outCh := make(chan int)

	go func() {
		for _, val := range unity {

			outCh <- val
		}

		close(outCh)
	}()

	return outCh
}

func sendValStr(unity []string) (ch <-chan string) {

	outCh := make(chan string)

	go func() {
		for _, val := range unity {

			outCh <- val
		}

		close(outCh)
	}()

	return outCh
}

func Test_channel_merge_01(t *testing.T) {

	_, err := MergeChanels()
	if err == nil {
		t.Errorf("it should raise error")
	}

}

func Test_channel_merge_02(t *testing.T) {

	data1 := []int{1, 2, 3, 4}
	ch := sendValInt(data1)
	result, err := MergeChanels(ch)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	dataOut := make([]int, 0)
	for out := range result.(<-chan int) {
		dataOut = append(dataOut, out)
	}

	eq, err := NewCollection(data1).IsElemEqual(dataOut)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	if !eq {
		t.Errorf("it should be equal on element\n")
	}

}

func Test_channel_merge_03(t *testing.T) {

	data1 := []string{"ani", "shinta", "ari"}
	ch := sendValStr(data1)
	result, err := MergeChanels(ch)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	dataOut := make([]string, 0)
	for out := range result.(<-chan string) {
		dataOut = append(dataOut, out)
	}

	eq, err := NewCollection(data1).IsElemEqual(dataOut)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	if !eq {
		t.Errorf("it should be equal on element\n")
	}

}

func Test_channel_merge_04(t *testing.T) {

	data1 := []int{1, 2, 3, 4}
	data2 := []int{5, 6}
	data3 := []int{7, 6, 9}

	ch1 := sendValInt(data1)
	ch2 := sendValInt(data2)
	ch3 := sendValInt(data3)
	result, err := MergeChanels(ch1, ch2, ch3)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	dataOut := make([]int, 0)
	for out := range result.(<-chan int) {
		dataOut = append(dataOut, out)
	}

	eq, err := NewCollection([]int{1, 2, 3, 4, 5, 6, 7, 6, 9}).IsElemEqual(dataOut)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}

	if !eq {
		t.Errorf("it should be equal on element\n")
	}

}
