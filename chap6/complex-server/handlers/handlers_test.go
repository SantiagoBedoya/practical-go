package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SantiagoBedoya/practical-go/chap6/complex-server/config"
)

func Test_apiHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	b := new(bytes.Buffer)
	c := config.InitConfig(b)

	apiHandler(w, r, c)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected response status %v, got %v", http.StatusOK, resp.StatusCode)
	}
	expectedResponseBody := "Hello World"
	if string(body) != expectedResponseBody {
		t.Errorf("expected response %s, got %s", expectedResponseBody, string(body))
	}
}

func Test_healthCheckHandler(t *testing.T) {
	tests := []struct {
		name   string
		method string
		output string
		status int
	}{
		{
			name:   "post method not allowed",
			method: "POST",
			output: "Method not allowed\n",
			status: http.StatusMethodNotAllowed,
		},
	}

	b := new(bytes.Buffer)
	c := config.InitConfig(b)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/health", nil)
			w := httptest.NewRecorder()
			healthCheckHandler(w, r, c)
			resp := w.Result()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error reading response body: %v", err)
			}
			if resp.StatusCode != tc.status {
				t.Errorf("expected response status %v, got %v", tc.status, resp.StatusCode)
			}
			if string(body) != tc.output {
				t.Errorf("expected response '%s', got '%s'", tc.output, string(body))
			}
		})
		b.Reset()
	}

}
