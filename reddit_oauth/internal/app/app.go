package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SantiagoBedoya/reddit_oauth/internal/config"
	"github.com/SantiagoBedoya/reddit_oauth/internal/handlers"
	"github.com/SantiagoBedoya/reddit_oauth/internal/repositories"
	"github.com/SantiagoBedoya/reddit_oauth/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	cfg := config.LoadConfig(".config.cfg")
	repo, err := repositories.NewMongoRepo(cfg.MongoURI, cfg.MongoDB, cfg.MongoCollection)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	notificatorSvc, err := service.NewNotificatorService(cfg)
	if err != nil {
		logrus.Error("notificator connection failed")
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
	})
	app.Use(limiter.New())
	app.Use(logger.New())

	svc := service.NewService(repo, notificatorSvc)
	handler := handlers.NewHandler(svc)
	app.Post("/api/v1/register", handler.Register)
	app.Post("/api/v1/login", handler.Login)

	errs := make(chan error, 2)
	go func() {
		errs <- app.Listen(fmt.Sprintf(":%d", cfg.Port))
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	<-errs
	logrus.Info("shutting down...")
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
