package cmd

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

type httpConfig struct {
	url  string
	verb string
}

func makeRequest(c httpConfig) ([]byte, error) {
	request, err := http.NewRequest(c.verb, c.url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	r, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
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

	switch c.verb {
	case "GET", "get", "POST", "post", "HEAD", "head":
		body, err := makeRequest(c)
		if err != nil {
			return err
		}
		fmt.Fprintln(w, string(body))
		return nil
	default:
		return ErrNoAllowedMethod
	}
}
