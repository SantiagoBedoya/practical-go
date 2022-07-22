package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func startBadTestHTTPServerV2(shutdownServer chan struct{}) *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				<-shutdownServer
				fmt.Fprintf(w, "Hello, World")
			},
		),
	)
	return ts
}

func Test_fetchRemoteResourceV2(t *testing.T) {
	shutdownServer := make(chan struct{})
	ts := startBadTestHTTPServerV2(shutdownServer)
	defer ts.Close()
	defer func() {
		shutdownServer <- struct{}{}
	}()
	client := createHTTPClientWithTimeout(200 * time.Millisecond)
	_, err := fetchRemoteResource(client, ts.URL)
	if err == nil {
		t.Logf("expected non-nil error")
		t.Fail()
	}
	if !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("expected error to contain: context deadline exceeded, got %v", err)
	}
}
