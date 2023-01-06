package service

import "errors"

var (
	ErrEmailInUse         = errors.New("email already in use")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
