package httperrors

import (
	"log"
	"net/http"
)

// HTTPError define data struct
type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

// NewHTTPError creates and return an instace of HTTPError
func NewHTTPError(statusCode int, msg, err string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    msg,
		Error:      err,
	}
}

// NewInternalServerError return HTTPError with InternalServerError info
func NewInternalServerError(msg string, err error) *HTTPError {
	log.Println(err.Error())
	return NewHTTPError(http.StatusInternalServerError, msg, "InternalServerError")
}

// NewUnexpectedError return HTTPError with InternalServerError info and default message
func NewUnexpectedError(err error) *HTTPError {
	return NewInternalServerError("something went wrong!", err)
}

// NewNotFoundError return HTTPError with NotFound info
func NewNotFoundError(msg string) *HTTPError {
	return NewHTTPError(http.StatusNotFound, msg, "NotFound")
}

// NewUnauthorizedError return HTTPError with Unauthorized info
func NewUnauthorizedError(msg string) *HTTPError {
	return NewHTTPError(http.StatusUnauthorized, msg, "Unauthorized")
}

// NewBadRequestError return HTTPError with BadRequest info
func NewBadRequestError(msg string) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, msg, "BadRequest")
}
