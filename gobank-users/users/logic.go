package users

import (
	"github.com/SantiagoBedoya/gobank-utils/httperrors"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

// NewService return the implementation of Service interface
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(data *User) (*User, *httperrors.HTTPError) {
	if err := data.ValidateLogin(); err != nil {
		return nil, err
	}
	u, err := s.repo.FindByEmail(data)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, httperrors.NewUnauthorizedError(ErrInvalidCredentials.Error())
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password)); err != nil {
		return nil, httperrors.NewUnauthorizedError(ErrInvalidCredentials.Error())
	}
	return u, nil
}

func (s *service) Create(data *User) (*User, *httperrors.HTTPError) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, httperrors.NewUnexpectedError(err)
	}
	data.Password = string(hash)
	u, err := s.repo.Create(data)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			return nil, httperrors.NewBadRequestError(ErrUserAlreadyExists.Error())
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	return u, nil
}

func (s *service) FindByID(userID string) (*User, *httperrors.HTTPError) {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, httperrors.NewNotFoundError(ErrUserNotFound.Error())
		}
		return nil, httperrors.NewUnexpectedError(err)
	}
	return u, nil
}
