package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	data string
	t    *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	output := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.data {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		output <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-output:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("Not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("Store data should be written to the body", func(t *testing.T) {
		data := "Hello world!"
		store := &SpyStore{data: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Want %q got %q", data, response.Body.String())
		}
	})

	t.Run("Should cancel if cancel called", func(t *testing.T) {
		data := "Hello World!"
		store := &SpyStore{data: data}

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Errorf("A response should not have been written")
		}
	})
}
