package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello James!"
	got := hello("James")

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestRscHello(t *testing.T) {
	want := "Hello, world."
	got := rscHello()

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
