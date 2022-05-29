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
	newYear2 := time.Date(2000, time.January, 3, 0, 0, 0, 0, time.Local)
	normal := time.Date(2000, time.December, 1, 0, 0, 0, 0, time.Local)

	cases := map[string]TestCase{
		"now": {nil},
		"newYear": {&newYear},
		"newYear2": {&newYear2},
		"normal": {&normal},
	}

	for _, v := range cases {
		testRun(t, v)
	}
}

// 質問したいこと
// - test.helperの重要性
// - bodyからfortuneだけを抽出する方法

func testRun(t *testing.T, tc TestCase) {
	// t.Helper()

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
