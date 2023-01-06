package service

import (
	"context"

	"github.com/SantiagoBedoya/reddit_oauth/internal/models"
)

type Service interface {
	Register(doc *models.User) error
	Login(doc *models.User) (*string, error)
}

type Repository interface {
	Create(ctx context.Context, doc *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
