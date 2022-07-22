package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startHTTPServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
		fmt.Fprintf(w, "I am the request header echoing program")
	}))
	return ts
}

func TestAddHeaderMiddleware(t *testing.T) {
	testsHeader := map[string]string{
		"X-Client-Id": "test-client",
		"X-Auth-Hash": "randomString",
	}
	client := createClient(testsHeader)
	ts := startHTTPServer()
	defer ts.Close()
	resp, err := client.Get(ts.URL)
	if err != nil {
		t.Fatalf("expected non-nil, got %v", err)
	}
	for k, v := range testsHeader {
		if resp.Header.Get(k) != testsHeader[k] {
			t.Fatalf("expected header %s:%s, got %s:%s", k, v, k, testsHeader[k])
		}
	}
}
