package users

import (
	"strings"

	"github.com/SantiagoBedoya/gobank-utils/httperrors"
)

// User define data struct
type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

// ValidateLogin make validations to data
func (u *User) ValidateLogin() *httperrors.HTTPError {
	if strings.TrimSpace(u.Email) == "" {
		return httperrors.NewBadRequestError("email should not be empty")
	}
	if strings.TrimSpace(u.Password) == "" {
		return httperrors.NewBadRequestError("password should not be empty")
	}
	return nil
}

// Validate make validations to data
func (u *User) Validate() *httperrors.HTTPError {
	if strings.TrimSpace(u.FirstName) == "" {
		return httperrors.NewBadRequestError("first name should not be empty")
	}
	if strings.TrimSpace(u.LastName) == "" {
		return httperrors.NewBadRequestError("last name should not be empty")
	}
	if strings.TrimSpace(u.Email) == "" {
		return httperrors.NewBadRequestError("email should not be empty")
	}
	if strings.TrimSpace(u.Password) == "" {
		return httperrors.NewBadRequestError("password should not be empty")
	}
	return nil
}
