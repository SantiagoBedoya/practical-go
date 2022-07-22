package client

import "net/http"

// AddHeaderMiddleware define data struct
type AddHeaderMiddleware struct {
	headers map[string]string
}

// RoundTrip define method
func (h AddHeaderMiddleware) RoundTrip(r *http.Request) (*http.Response, error) {
	reqCopy := r.Clone(r.Context())
	for k, v := range h.headers {
		reqCopy.Header.Add(k, v)
	}
	return http.DefaultTransport.RoundTrip(reqCopy)
}

func createClient(headers map[string]string) *http.Client {
	h := AddHeaderMiddleware{
		headers: headers,
	}
	client := http.Client{
		Transport: &h,
	}
	return &client
}
