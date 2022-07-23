package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_setupServer(t *testing.T) {
	b := new(bytes.Buffer)
	mux := http.NewServeMux()
	wrappedMux := setupServer(mux, b)
	ts := httptest.NewServer(wrappedMux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/panic")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf(
			"expected response status %v, got %v",
			http.StatusInternalServerError,
			resp.StatusCode,
		)
	}
}
