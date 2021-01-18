package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "James")

	got := buffer.String()
	want := "Hello James"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
