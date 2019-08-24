package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Index(response, request)

	t.Run("returns index page", func(t *testing.T) {
		got := response.Body.String()
		want := "pair monitor"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}
