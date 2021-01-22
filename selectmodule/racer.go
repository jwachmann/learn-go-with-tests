package selectmodule

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

// Racer checks each of the given urls to see which one responds the fastest, will timeout if neither return within 10 seconds
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer is a racer with a configurable timeour
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
