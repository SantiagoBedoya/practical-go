package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// LoggingClient define data struct
type LoggingClient struct {
	log *log.Logger
}

// RoundTrip define RoudTrip to provide method
func (c LoggingClient) RoundTrip(r *http.Request) (*http.Response, error) {
	c.log.Printf("Sending a %s request to %s over %s", r.Method, r.URL, r.Proto)
	resp, err := http.DefaultTransport.RoundTrip(r)
	c.log.Printf("Got back a response over %s\n", resp.Proto)
	return resp, err
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{Timeout: d}
}

func createHTTPGetRequestWithContext(ctx context.Context, url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return req, err
}

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "must specify a HTTP URL to get data from")
		os.Exit(1)
	}
	myTransport := LoggingClient{}
	l := log.New(os.Stdout, "", log.LstdFlags)
	myTransport.log = l

	client := createHTTPClientWithTimeout(15 * time.Second)
	client.Transport = &myTransport

	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%#v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Bytes in response: %d\n", len(body))
}
