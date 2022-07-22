package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
)

func createHTTPGetRequestWithTrace(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	trace := &httptrace.ClientTrace{
		DNSDone: func(di httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", di)
		},
		GotConn: func(gci httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", gci)
		},
	}
	ctxTrace := httptrace.WithClientTrace(req.Context(), trace)
	req.WithContext(ctxTrace)
	return req, err
}

func createHTTPClientWithTimeout(d time.Duration) *http.Client {
	return &http.Client{Timeout: d}
}

func main() {
	d := 5 * time.Second
	ctx := context.Background()
	client := createHTTPClientWithTimeout(d)

	req, err := createHTTPGetRequestWithTrace(ctx, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		client.Do(req)
		time.Sleep(1 * time.Second)
		fmt.Println("-----------")
	}
}
