package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SantiagoBedoya/gobank-users-api/api"
	"github.com/SantiagoBedoya/gobank-users-api/repositories/postgres"
	"github.com/SantiagoBedoya/gobank-users-api/users"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("error loading .env file")
// 	}
// }

func main() {
	datasourceName, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		log.Fatal("POSTGRES_URL is not defined")
	}
	timeout := time.Second * 5
	port := ":8080"
	db, err := sql.Open("pgx", datasourceName)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	defer db.Close()

	app := gin.Default()
	app.Use(cors.Default())
	app.Use(helmet.Default())
	// app.Use(gzip.Gzip(gzip.DefaultCompression))
	r := app.Group("/api/v1/users")
	{
		repository := postgres.NewPostgreSQLRepository(db, timeout)
		srv := users.NewService(repository)
		handler := api.NewHandler(srv)
		r.POST("", handler.Create)
		r.POST("/login", handler.Login)
		r.GET("/:id", handler.FindByID)
	}

	server := http.Server{
		Addr:         port,
		Handler:      app,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	// gracefull shutdown
	errs := make(chan error, 2)
	go func() {
		c := make(chan error, 1)
		errs <- fmt.Errorf("%s", <-c)
	}()
	go func() {
		log.Printf("Service running on port %s", port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("ERROR: %v\n", <-errs)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}
