package users

import "github.com/SantiagoBedoya/gobank-utils/httperrors"

// Service define interface for services
type Service interface {
	Create(*User) (*User, *httperrors.HTTPError)
	FindByID(string) (*User, *httperrors.HTTPError)
	Login(*User) (*User, *httperrors.HTTPError)
}
