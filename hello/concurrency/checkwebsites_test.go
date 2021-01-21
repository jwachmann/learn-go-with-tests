package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "www.verticalscope.com" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"www.google.com",
		"www.kobo.com",
		"www.verticalscope.com",
	}

	want := map[string]bool{
		urls[0]: true,
		urls[1]: true,
		urls[2]: false,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wanted '%v' but got '%v'", want, got)
	}
}

func slowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
