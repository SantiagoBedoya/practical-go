package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/SantiagoBedoya/auth-ms/internal/config"
	"github.com/SantiagoBedoya/auth-ms/internal/pb"
	"github.com/SantiagoBedoya/auth-ms/internal/repository"
	"github.com/SantiagoBedoya/auth-ms/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Initialize(configPath string) {
	cfg := config.LoadConfig(configPath)
	privateKey, err := ioutil.ReadFile(cfg.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err := ioutil.ReadFile(cfg.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.NewPostgresRepo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	svc := service.NewService(repo, privateKey, publicKey)
	pb.RegisterAuthServer(grpcServer, svc)
	reflection.Register(grpcServer)

	errs := make(chan error, 2)
	go func() {
		log.Printf("gRPC server is running on port %d", cfg.Port)
		errs <- grpcServer.Serve(listen)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	<-errs
	log.Println("Shutting down...")
	grpcServer.GracefulStop()
}
