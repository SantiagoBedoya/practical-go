package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var errInvalidSubCommand = errors.New("invalid sub-command specified")

func handleCmdA(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fmt.Fprintf(w, "executing command A")
	return nil
}

func handleCmdB(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	if err := fs.Parse(args); err != nil {
		return err
	}
	fmt.Fprintf(w, "executing command B")
	return nil
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [cmd-a|cmd-b] -h\n", os.Args[0])
	handleCmdA(w, []string{"-h"})
	handleCmdB(w, []string{"-h"})
}

func main() {
	var err error
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "cmd-a":
		err = handleCmdA(os.Stdout, os.Args[2:])
	case "cmd-b":
		err = handleCmdB(os.Stdout, os.Args[2:])
	default:
		printUsage(os.Stdout)
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}
