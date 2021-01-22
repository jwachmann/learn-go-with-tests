package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("Word doesn't exist", func(t *testing.T) {
		dictionary := Dictionary{}

		got, err := dictionary.Search("test")
		var want string

		assertError(t, err, ErrNotFound)
		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")

	got, err := dictionary.Search("test")
	want := "this is just a test"

	assertError(t, err, nil)
	assertStrings(t, got, want)
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	dictionary.Delete("test")

	got, err := dictionary.Search("test")
	var want string

	assertError(t, err, ErrNotFound)
	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error %q want %q", got, want)
	}
}
