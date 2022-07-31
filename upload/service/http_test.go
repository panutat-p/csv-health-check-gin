package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Status OK
func TestCheck(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ch := make(chan bool, 1)
	Check(ch, srv.URL)
	got := <-ch

	if !got {
		t.Error("want true")
	}
}

// Status NotFound, invalid URL
func TestCheck2(t *testing.T) {
	ch := make(chan bool, 1)
	Check(ch, "sss")
	got := <-ch

	if got {
		t.Error("want false")
	}
}

// HTTP client error
func TestCheck3(t *testing.T) {
	ch := make(chan bool, 1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		panic("unexpected")
	}))
	defer srv.Close()

	Check(ch, srv.URL)
	got := <-ch

	if got {
		t.Error("want false")
	}
}

// Request timeout
func TestCheck4(t *testing.T) {
	ch := make(chan bool, 1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestTimeout)
		time.Sleep(6 * time.Second)
	}))
	defer srv.Close()

	Check(ch, srv.URL)
	got := <-ch

	if got {
		t.Error("want false")
	}
}
