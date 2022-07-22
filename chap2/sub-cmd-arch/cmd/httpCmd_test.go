package cmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleHTTP(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		err    error
		output string
	}{
		{
			name: "server not specified",
			args: []string{""},
			err:  ErrNoServerSpecified,
		},
		{
			name: "help flag",
			args: []string{"-h"},
			err:  errors.New("flag: help requested"),
		},
		{
			name:   "passing server",
			args:   []string{"localhost"},
			err:    nil,
			output: "executing http command\n",
		},
		{
			name: "method not allowed",
			args: []string{"-verb", "PUT", "localhost"},
			err:  ErrNoAllowedMethod,
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := HandleHTTP(byteBuf, tc.args)
			if tc.err != nil && err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
			}
			if tc.err == nil && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if len(tc.output) != 0 {
				gotMsg := byteBuf.String()
				if tc.output != gotMsg {
					t.Errorf("expected output %v, got %v", tc.output, gotMsg)
				}
			}
			byteBuf.Reset()
		})
	}
}
