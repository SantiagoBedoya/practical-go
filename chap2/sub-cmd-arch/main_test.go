package main

import (
	"bytes"
	"testing"

	"github.com/SantiagoBedoya/practical-go/chap2/sub-cmd-arch/cmd"
)

func Test_handleCommand(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		output string
		err    error
	}{
		{
			name: "no sub-command specified",
			args: []string{},
			err:  errInvalidSubCommand,
		},
		{
			name: "help flag",
			args: []string{"-h"},
			err:  nil,
		},
		{
			name: "invalid sub-command",
			args: []string{"foo"},
			err:  errInvalidSubCommand,
		},
		{
			name: "calling http command without server",
			args: []string{"http"},
			err:  cmd.ErrNoServerSpecified,
		},
		{
			name: "calling grpc command without server",
			args: []string{"grpc"},
			err:  cmd.ErrNoServerSpecified,
		},
		{
			name:   "calling http command",
			args:   []string{"http", "http://localhost"},
			err:    nil,
			output: "executing http command\n",
		},
		{
			name:   "calling grpc command",
			args:   []string{"grpc", "http://localhost"},
			err:    nil,
			output: "executing grpc command\n",
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := handleCommand(byteBuf, tc.args)
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
					t.Errorf("expected output '%v', got '%v'", tc.output, gotMsg)
				}
			}
			byteBuf.Reset()
		})
	}
}
