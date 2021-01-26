package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing a counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("It runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		wg := sync.WaitGroup{}
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, counter *Counter, count int) {
	t.Helper()

	if counter.Value() != count {
		t.Errorf("Wanted '%v' got '%v'", count, counter.Value())
	}
}
