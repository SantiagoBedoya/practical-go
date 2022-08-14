package accounts

import "errors"

var (
	//ErrInvalidCredentials is the error when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	//ErrUserAlreadyExists is the error when the email is already in use
	ErrUserAlreadyExists = errors.New("user already exists")
)
