package simhelper_test

import (
	"time"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func ExampleWait() {

	simhelper.Wait(5 * time.Second) // sleep / wait for 5 seconds
}
