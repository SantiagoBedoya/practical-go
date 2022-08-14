package accounts

import "github.com/go-playground/validator/v10"

// Request define data struct for request
type Request struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Validate make validation to request data struct
func (r Request) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Response define data struct for response
type Response struct {
	StatusCode  int    `json:"status_code"`
	AccessToken string `json:"access_token,omitempty"`
	Error       string `json:"error,omitempty"`
}
