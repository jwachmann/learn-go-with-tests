package mocking

import (
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// Sleeper represents something that'll sleep (block) for a fixed duration in-between counts while counting down
type Sleeper interface {
	Sleep()
}

// Countdown counts down from 3 by printing to the given writer
func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}

	fmt.Fprint(writer, finalWord)
}
