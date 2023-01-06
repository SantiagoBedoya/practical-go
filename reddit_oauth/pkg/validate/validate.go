package validate

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type HTTPError struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

type ValidateError struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func Validate(data interface{}) *HTTPError {
	validate := validator.New()
	errs := make([]*ValidateError, 0)
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, &ValidateError{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
		return &HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
			Data:       errs,
		}
	} else {
		return nil
	}
}
