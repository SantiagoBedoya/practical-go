package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SantiagoBedoya/todo-app/config"
	"github.com/SantiagoBedoya/todo-app/internal/services"
	"github.com/SantiagoBedoya/todo-app/internal/transport/rest"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Run initialize the server
func Run() {
	cfg := config.LoadConfig("./config/config.cfg")
	redis := getRedisClient(cfg)

	app := fiber.New(fiber.Config{
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		IdleTimeout:           3 * time.Second,
		DisableStartupMessage: true,
	})

	service := services.NewTodoService(redis)
	rest.RegisterRoutes(app, service)

	errs := make(chan error, 2)
	go func() {
		port := fmt.Sprintf(":%d", cfg.Port)
		logrus.Info("Server is running on port ", port)
		errs <- app.Listen(port)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	<-errs
	logrus.Info("Shutting down...")
	app.Shutdown()
}

func getRedisClient(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password:     "",
		DB:           0,
		MaxIdleConns: 5,
	})
	return rdb
}
