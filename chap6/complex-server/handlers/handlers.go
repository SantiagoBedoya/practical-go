package handlers

import (
	"fmt"
	"net/http"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
)

type app struct {
	conf    config.AppConfig
	handler func(w http.ResponseWriter, r *http.Request, conf config.AppConfig)
}

func (a app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler(w, r, a.conf)
}

func apiHandler(w http.ResponseWriter, r *http.Request, config config.AppConfig) {
	config.Logger.Println("Handling API request")
	fmt.Fprintf(w, "Hello World")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, conf config.AppConfig) {
	if r.Method != http.MethodGet {
		conf.Logger.Printf("error=\"invalid request\" path=%s method=%s", r.URL.Path, r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	conf.Logger.Println("Handling healthcheck request")
	fmt.Fprintf(w, "ok")
}

func panicHandler(w http.ResponseWriter, r *http.Request, config config.AppConfig) {
	panic("I panicked")
}
