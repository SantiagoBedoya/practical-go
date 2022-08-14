package accounts

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// AuthService define blueprint for service
type AuthService interface {
	SignIn(context.Context, string, string) (string, error)
	SignUp(context.Context, string, string) (string, error)
}

// MakeSignInEndpoint manage request and response for service
func MakeSignInEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)
		if err := req.Validate(); err != nil {
			return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusBadRequest}, nil
		}
		token, err := svc.SignIn(ctx, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, ErrInvalidCredentials) {
				return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusUnauthorized}, nil
			}
			return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusInternalServerError}, nil
		}
		return Response{AccessToken: token, Error: "", StatusCode: http.StatusOK}, nil
	}
}

// MakeSignUpEndpoint manage request and response for service
func MakeSignUpEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)
		if err := req.Validate(); err != nil {
			return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusBadRequest}, nil
		}
		token, err := svc.SignUp(ctx, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, ErrUserAlreadyExists) {
				return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusBadRequest}, nil
			}
			return Response{AccessToken: "", Error: err.Error(), StatusCode: http.StatusInternalServerError}, nil
		}
		return Response{AccessToken: token, Error: "", StatusCode: http.StatusCreated}, nil
	}
}
