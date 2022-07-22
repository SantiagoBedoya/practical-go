package cmd

import "errors"

// ErrNoServerSpecified is when server is no specified
var ErrNoServerSpecified = errors.New("you must to specify a server")
