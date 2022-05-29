package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestCase struct {
	time *time.Time
}

func TestHomeHandler(t *testing.T) {
	newYear := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)
	normal := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local)
	cases := map[string]TestCase{
		"test1": {nil},
		"test2": {&newYear},
		"test3": {&normal},
	}
	for _, v := range cases {
		testRun(t, v)
	}
}

func testRun(t *testing.T, tc TestCase) {
	t.Helper()
	testServer := httptest.NewServer(http.HandlerFunc(makeHandler(homeHandler, tc.time)))
	defer testServer.Close()

	res, err := http.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	_, err = io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}
}
