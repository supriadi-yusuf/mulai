package simhelper

import "context"

// CheckForCancelation is function to check if current process is being canceled outside
func CheckForCancelation(ctx context.Context) {
	select {
	case <-ctx.Done():
		panic(ctx.Err().Error())
	default:
	}
}

// DoCancelRecover is function to stop panic propagation
func DoCancelRecover() {
	recover()
}
