package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(homeHandler))
	defer testServer.Close()

	res, err := http.Get(testServer.URL)
	if err != nil {
			t.Error(err)
	}

	_, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
			t.Error(err)
	}

	if res.StatusCode != 200 {
			t.Error("a response code is not 200")
	}
}
