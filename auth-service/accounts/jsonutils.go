package accounts

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodeRequest decode request to JSON
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeResponse encode response to JSON
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	res, ok := response.(Response)
	if ok {
		w.WriteHeader(res.StatusCode)
	}
	return json.NewEncoder(w).Encode(response)
}
