package simhelper_test

import (
	"context"
	"time"

	"github.com/supriadi-yusuf/mulai/simhelper"
)

func ExampleCheckForCancelation() {

	var maxCntr = 10000
	var cntr int = 0

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // do cancelation

	go func() {

		defer simhelper.DoCancelRecover()

		for cntr = 0; cntr < maxCntr; cntr++ {
			simhelper.Wait(10 * time.Millisecond)
			simhelper.CheckForCancelation(ctx)
		}
	}()

	simhelper.Wait(1000 * time.Millisecond)

}
