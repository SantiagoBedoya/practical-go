package main

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
		config
	}{
		{
			name:   "help flag",
			args:   []string{"-h"},
			err:    errors.New("flag: help requested"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			name:   "print 10 times",
			args:   []string{"-n", "10"},
			err:    errors.New("positional arguments specified"),
			config: config{printUsage: false, numTimes: 10},
		},
		{
			name:   "invalid args",
			args:   []string{"-n", "abc"},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			name:   "invalid number of arguments",
			args:   []string{"-n", "1", "foo"},
			err:    errors.New("positional arguments specified"),
			config: config{printUsage: false, numTimes: 1},
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c, err := parseArgs(byteBuf, tc.args)
			if tc.err == nil && err != nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if c.printUsage != tc.printUsage {
				t.Errorf("expected printUsage %v, got %v", tc.printUsage, c.printUsage)
			}
			if c.numTimes != tc.numTimes {
				t.Errorf("expected numTimes %v, got %v", tc.numTimes, c.numTimes)
			}
		})
	}
}

func Test_validateArgs(t *testing.T) {
	tests := []struct {
		name string
		c    config
		err  error
	}{
		{
			name: "empty config",
			c:    config{},
			err:  errors.New("must specify a number greater than 0"),
		},
		{
			name: "invalid numTimes",
			c:    config{numTimes: -1},
			err:  errors.New("must specify a number greater than 0"),
		},
		{
			name: "valid numtimes",
			c:    config{numTimes: 10},
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateArgs(tc.c)
			if tc.err != nil && err == nil {
				t.Errorf("expetected error %v, got %v", tc.err, err)
			}
			if tc.err != nil && err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expetected error message %v, got %v", tc.err, err)
				}
			}
		})
	}
}

func Test_runCmd(t *testing.T) {
	tests := []struct {
		name   string
		c      config
		input  string
		output string
		err    error
	}{
		{
			name:   "printUsage",
			c:      config{printUsage: true},
			output: usageString,
		},
		{
			name:   "empty input",
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("Your name please? Press the enter key when done.\n", 1),
			err:    errors.New("you didn't enter your name"),
		},
		{
			name:   "correct way",
			c:      config{numTimes: 5},
			input:  "Santiago",
			output: "Your name please? Press the enter key when done.\n" + strings.Repeat("Nice to meet you Santiago\n", 5),
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rd := strings.NewReader(tc.input)
			err := runCmd(rd, byteBuf, tc.c)
			if err != nil && tc.err == nil {
				t.Errorf("expected nil error, got %v", err)
			}
			if tc.err != nil && err != nil {
				if tc.err.Error() != err.Error() {
					t.Errorf("expected error message %v, got %v", tc.err, err)
				}
			}
			gotMsg := byteBuf.String()
			if gotMsg != tc.output {
				t.Errorf("expected output %v, got %v", tc.output, gotMsg)
			}
			byteBuf.Reset()
		})
	}
}

func TestMain(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "execute without args",
			args: []string{},
			err:  errors.New("exit status 1"),
		},
		{
			name: "execute with help flag",
			args: []string{"-h"},
			err:  errors.New("exit status 1"),
		},
		{
			name: "execute with numTimes arg",
			args: []string{"5", "Santiago"},
			err:  errors.New("exit status 1"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := exec.Command("./app", tc.args...).Run()
			if tc.err == nil && err != nil {
				t.Errorf("expected nil err, got %v", err)
			}
			if err != nil && tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("expected error message %v, got %v", tc.err, err)
				}
			}
		})
	}
}
