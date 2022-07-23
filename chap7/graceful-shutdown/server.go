package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func shutdown(ctx context.Context, s *http.Server, waitForShutdownCompletion chan struct{}) {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigch
	log.Printf("Got signal %v. Server shutting down.", sig)
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("error during shutdown %v", err)
	}
	waitForShutdownCompletion <- struct{}{}
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("I started processing the request")
	time.Sleep(35 * time.Second)
	log.Println("Before continuing, i will check if the timeout has already expired")
	if r.Context().Err() != nil {
		log.Printf("Aborting further processing %v\n", r.Context().Err())
	}
	fmt.Fprintf(w, "Hello World")
	log.Println("I finished processing the request")
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	waitForShutdownCompletion := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/", handleUserAPI)

	s := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}
	go shutdown(ctx, &s, waitForShutdownCompletion)
	err := s.ListenAndServe()
	log.Print("Waiting for shutdown to complete...")
	<-waitForShutdownCompletion
	log.Fatal(err)

}
