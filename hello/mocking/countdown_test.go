package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyWriterSleeper struct {
	Calls []string
}

func (s *SpyWriterSleeper) Write(bytes []byte) (int, error) {
	s.Calls = append(s.Calls, write)
	return 0, nil
}

func (s *SpyWriterSleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func TestCountdown(t *testing.T) {
	t.Run("Should print the correct values while counting down", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyWriterSleeper{}
		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	})

	t.Run("Should pause in-between printing the count", func(t *testing.T) {
		spy := &SpyWriterSleeper{}
		Countdown(spy, spy)

		got := spy.Calls
		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Did not sleep correctly in-between counts, got '%v' wanted '%v'", got, want)
		}
	})
}
