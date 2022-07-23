package handlers

import (
	"net/http"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
)

// Register define path and handler
func Register(mux *http.ServeMux, conf config.AppConfig) {
	mux.Handle("/health", &app{conf: conf, handler: healthCheckHandler})
	mux.Handle("/api", &app{conf: conf, handler: apiHandler})
	mux.Handle("/panic", &app{conf: conf, handler: panicHandler})
}
