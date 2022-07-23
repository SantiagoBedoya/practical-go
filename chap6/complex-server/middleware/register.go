package middleware

import (
	"net/http"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
)

// RegisterMiddleware attach middlewares to server
func RegisterMiddleware(mux *http.ServeMux, c config.AppConfig) http.Handler {
	return loggingMiddleware(panicMiddleware(mux, c), c)
}
