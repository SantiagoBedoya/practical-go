package middleware

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/handlers"
)

func Test_panicMiddleware(t *testing.T) {
	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	m := http.NewServeMux()
	handlers.Register(m, c)

	h := panicMiddleware(m, c)
	r := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	resp := w.Result()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body %v", err)
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf(
			"expected response status: %v, got %v",
			http.StatusOK,
			http.StatusInternalServerError,
		)
	}
	expectedResponseBody := "Unexpected server error"
	if string(body) != expectedResponseBody {
		t.Errorf(
			"expected response: %s, got %s",
			expectedResponseBody,
			string(body),
		)
	}
}
