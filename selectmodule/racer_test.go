package selectmodule

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()

	got := Racer(slowServer.URL, fastServer.URL)
	want := fastServer.URL

	if want != got {
		t.Errorf("Want %q got %q", want, got)
	}
}