package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/SantiagoBedoya/practical-go/chap2/sub-cmd-arch/cmd"
)

var errInvalidSubCommand = errors.New("invalid sub-command specified")

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			err = cmd.HandleHTTP(w, args[1:])
		case "grpc":
			err = cmd.HandleGRPC(w, args[1:])
		case "-h", "-help", "--help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}
	if errors.Is(err, cmd.ErrNoServerSpecified) || errors.Is(err, errInvalidSubCommand) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [cmd-a|cmd-b] -h\n", os.Args[0])
	cmd.HandleHTTP(w, []string{"-h"})
	cmd.HandleGRPC(w, []string{"-h"})
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
