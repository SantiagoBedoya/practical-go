package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcConfig struct {
	server string
	method string
	body   string
}

// HandleGRPC handle grpc command
func HandleGRPC(w io.Writer, args []string) error {
	c := grpcConfig{}
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.method, "method", "", "method to call")
	fs.StringVar(&c.body, "body", "", "body of request")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}
	c.server = fs.Arg(0)
	fmt.Fprintln(w, "executing grpc command")
	return nil
}
