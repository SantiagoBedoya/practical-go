package service

import (
	"context"
	"fmt"

	"github.com/SantiagoBedoya/reddit_notificator/internal/pb"
)

type service struct {
}

func NewService() pb.NotificatorServiceServer {
	return &service{}
}

func (s *service) Create(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	fmt.Println(req.Email, req.Content)
	return &pb.NotificationResponse{Done: true}, nil
}
