package service

import "errors"

var (
	ErrEmailUsernameInUse = errors.New("email/username already in use")
)
