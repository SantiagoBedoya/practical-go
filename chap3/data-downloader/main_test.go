package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, World")
			},
		),
	)
	return ts
}

func Test_fetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()
	expected := "Hello, World"
	data, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(data) {
		t.Errorf("expected response to be %v, got %v", expected, string(data))
	}
}
