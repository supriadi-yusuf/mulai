package simhelper

import (
	"log"
	"testing"
	"time"
)

func Test_Wait_01(t *testing.T) {

	log.Println(t.Name())

	d1 := 5 * time.Second
	startTime := time.Now()
	Wait(d1)
	d2 := time.Since(startTime)

	fsecond1 := d1.Milliseconds()
	fsecond2 := d2.Milliseconds()

	//fmt.Printf("s1 : %d\n", fsecond1)
	//fmt.Printf("s2 : %d\n", fsecond2)

	dseconds := fsecond2 - fsecond1
	if dseconds > 100 {
		t.Errorf("wait too long")
	}
}
