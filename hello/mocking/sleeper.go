package mocking

import "time"

// ConfigurableSleeper is a sleeper where the sleep duration and sleep implementation can be overridden
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// NewConfigurableSleeper creates a new instance configured with the given parameters
func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{
		duration,
		sleep,
	}
}

// Sleep sleeps for the time configured under "ConfigurableSleeper.duration"
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
