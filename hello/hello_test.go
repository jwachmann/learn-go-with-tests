package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		want := "Hello James!"
		got := hello("James", "")
		assertMatch(t, got, want)
	})

	t.Run("Say 'Hello World!' when an empty string is supplied", func(t *testing.T) {
		want := "Hello World!"
		got := hello("", "")
		assertMatch(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		want := "Hola Dora!"
		got := hello("Dora", "Spanish")
		assertMatch(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		want := "Bonjour Phileas!"
		got := hello("Phileas", "French")
		assertMatch(t, got, want)
	})
}

func TestRscHello(t *testing.T) {
	want := "Hello, world."
	got := rscHello()
	assertMatch(t, got, want)
}

func assertMatch(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
