package panicutils

import "log"

func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v", r)
	}
}

// RunWithRecovery Run a function with panic recovery
func RunWithRecovery(fn func()) {
	defer recoverFromPanic()
	fn()
}

// RunWithGoroutineRecovery Run a function with panic recovery in a goroutine
func RunWithGoroutineRecovery(fn func()) {
	go func() {
		defer recoverFromPanic()
		fn()
	}()
}
