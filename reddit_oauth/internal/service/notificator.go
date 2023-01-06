package service

import (
	"context"
	"fmt"
	"time"

	"github.com/SantiagoBedoya/reddit_oauth/internal/config"
	"github.com/SantiagoBedoya/reddit_oauth/internal/pb"
	"google.golang.org/grpc"
)

func NewNotificatorService(cfg *config.Config) (pb.NotificatorServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%d", cfg.NotificatorHost, cfg.NotificatorPort),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)

	if err != nil {
		return nil, err
	}
	client := pb.NewNotificatorServiceClient(conn)
	return client, nil
}
