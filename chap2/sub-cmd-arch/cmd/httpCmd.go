package cmd

import (
	"flag"
	"fmt"
	"io"
)

type httpConfig struct {
	url  string
	verb string
}

// HandleHTTP handle command http
func HandleHTTP(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "GET", "HTTP method")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}
	c := httpConfig{verb: v}
	c.url = fs.Arg(0)
	fmt.Fprintln(w, "executing http command")
	return nil
}
