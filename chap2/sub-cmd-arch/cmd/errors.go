package cmd

import "errors"

// ErrNoServerSpecified is when server is no specified
var ErrNoServerSpecified = errors.New("you must to specify a server")

// ErrNoAllowedMethod is when provided method is not allowed
var ErrNoAllowedMethod = errors.New("method not allowed")
