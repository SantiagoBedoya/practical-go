package main

import (
	"context"
	"testing"
	"time"
)

func Test_getNameContext(t *testing.T) {
	ctx1, cancel := context.WithTimeout(context.Background(), 5*time.Nanosecond)
	defer cancel()
	tests := []struct {
		name   string
		ctx    context.Context
		output string
		err    error
	}{
		{
			name:   "context deadline",
			ctx:    ctx1,
			output: "Default Name",
			err:    context.DeadlineExceeded,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			name, err := getNameContext(ctx1)
			if tc.err == nil && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if tc.err != nil && err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expected error '%v', got '%v'", tc.err, err)
				}
			}
			if tc.output != name {
				t.Errorf("expected output %v, got %v", tc.output, name)
			}
		})
	}
}
