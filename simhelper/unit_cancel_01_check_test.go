package simhelper

import (
	"context"
	"log"
	"testing"
	"time"
)

func Test_cancel_check_01(t *testing.T) {

	log.Println(t.Name())

	var maxCntr = 10000
	var cntr int = 0

	defer func() {

		if cntr >= maxCntr {
			t.Errorf("error")
		}

	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {

		defer DoCancelRecover()

		for cntr = 0; cntr < maxCntr; cntr++ {
			Wait(10 * time.Millisecond)
			CheckForCancelation(ctx)
		}
	}()

	Wait(1000 * time.Millisecond)
}
