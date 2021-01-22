package mocking

import (
	"testing"
	"time"
)

type SpyConfigurableSleeper struct {
	timeSlept time.Duration
}

func (s *SpyConfigurableSleeper) Sleep(duration time.Duration) {
	s.timeSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 1 * time.Second
	spyConfigurableSleeper := &SpyConfigurableSleeper{}

	sleeper := &ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyConfigurableSleeper.Sleep,
	}

	sleeper.Sleep()

	got := spyConfigurableSleeper.timeSlept
	want := sleepTime

	if got != want {
		t.Errorf("Expected to sleep for '%v' but actually slept for '%v'", want, got)
	}
}
