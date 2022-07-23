package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func progressStreamer(logReader *io.PipeReader, w http.ResponseWriter, done chan struct{}) {
	buf := make([]byte, 500)
	f, flushSupported := w.(http.Flusher)
	defer logReader.Close()
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Context-Type-Options", "nosniff")
	for {
		n, err := logReader.Read(buf)
		if err == io.EOF {
			break
		}
		w.Write(buf[:n])
		if flushSupported {
			f.Flush()
		}
	}
	done <- struct{}{}
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("I started processing the request")
	defer r.Body.Close()
	done := make(chan struct{})
	logReader, logWriter := io.Pipe()
	go longRunningProcess(logWriter)
	go progressStreamer(logReader, w, done)
	<-done
}

func longRunningProcess(w *io.PipeWriter) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "Hello")
		time.Sleep(1 * time.Second)
	}
	w.Close()
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/", handleUserAPI)
	s := http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
