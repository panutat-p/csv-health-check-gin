package upload

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Status OK
func TestCheck(t *testing.T) {
	ch := make(chan bool, 1)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer srv.Close()

	Check(ch, srv.URL)
	got := <-ch

	if !got {
		t.Error("want true")
	}
}

// Status NotFound
func TestCheck2(t *testing.T) {
	ch := make(chan bool, 1)

	Check(ch, "")
	got := <-ch

	if got {
		t.Error("want false")
	}
}

// unexpected error
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
