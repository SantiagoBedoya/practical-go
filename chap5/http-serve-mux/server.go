package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthCheckHandler)
	mux.HandleFunc("/api", apiHandler)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	mux := http.NewServeMux()
	setupHandlers(mux)
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
