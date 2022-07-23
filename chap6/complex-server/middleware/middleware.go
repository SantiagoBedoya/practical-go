package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
)

func loggingMiddleware(h http.Handler, c config.AppConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			h.ServeHTTP(w, r)
			log.Printf("path=%s method=%s duration=%f", r.URL.Path, r.Method, time.Now().Sub(startTime).Seconds())
		},
	)
}

func panicMiddleware(h http.Handler, c config.AppConfig) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rValue := recover(); rValue != nil {
					c.Logger.Println("panic detected", rValue)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Unexpected server error")
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}
