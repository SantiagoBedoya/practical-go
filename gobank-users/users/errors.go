package users

import "errors"

var (
	// ErrInvalidCredentials error when user credentials are invalid
	ErrInvalidCredentials = errors.New("credential are not valid")
	// ErrUserNotFound error when user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists error when user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
)
