package service

import (
	"context"
	"time"

	"github.com/SantiagoBedoya/reddit_oauth/internal/models"
	"github.com/SantiagoBedoya/reddit_oauth/internal/pb"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	notificationSvc pb.NotificatorServiceClient
	repo            Repository
}

func NewService(repo Repository, notificationSvc pb.NotificatorServiceClient) Service {
	return &service{notificationSvc, repo}
}

func (s service) Register(doc *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(doc.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	doc.Password = string(hash)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = s.repo.Create(ctx, doc); err != nil {
		return err
	}
	res, err := s.notificationSvc.Create(ctx, &pb.NotificationRequest{
		Email:   doc.Email,
		Content: "Welcome to the app",
	})
	if err != nil {
		logrus.Error("notificator service error", err)
	}
	logrus.Info("notificator response: ", res.Done)
	return nil
}

func (s service) Login(doc *models.User) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	user, err := s.repo.FindByEmail(ctx, doc.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(doc.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	uid := uuid.NewString()
	return &uid, nil
}
