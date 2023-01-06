package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/SantiagoBedoya/reddit_notificator/internal/config"
	"github.com/SantiagoBedoya/reddit_notificator/internal/pb"
	"github.com/SantiagoBedoya/reddit_notificator/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Initialize() {
	cfg := config.LoadConfig(".config.cfg")

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	svc := service.NewService()
	pb.RegisterNotificatorServiceServer(grpcServer, svc)

	errs := make(chan error, 2)
	go func() {
		logrus.Info("GRPC server listen on ", cfg.Port)
		errs <- grpcServer.Serve(listen)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	<-errs
	logrus.Info("shutting down...")
	grpcServer.GracefulStop()
}
