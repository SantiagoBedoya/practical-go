package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SantiagoBedoya/auth-service/accounts"
	"github.com/SantiagoBedoya/auth-service/repositories/postgresql"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	repository := postgresql.NewRepository(ctx, os.Getenv("DB_URI"), logger)
	defer repository.CloseConnection(ctx)

	var svc accounts.AuthService
	svc = accounts.AuthServiceInstance{
		Repository: repository,
	}
	svc = accounts.LogginMiddleware{Logger: logger, Next: svc}

	signInHandler := httptransport.NewServer(accounts.MakeSignInEndpoint(svc), accounts.DecodeRequest, accounts.EncodeResponse)
	signUpHandler := httptransport.NewServer(accounts.MakeSignUpEndpoint(svc), accounts.DecodeRequest, accounts.EncodeResponse)

	mux := mux.NewRouter()
	mux.Handle("/sign-in", signInHandler).Methods("POST")
	mux.Handle("/sign-up", signUpHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT MUST BE SET")
	}
	server := http.Server{
		Handler:      mux,
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	errs := make(chan error, 2)
	go func() {
		logger.Log("status", "ready", "addr", port)
		errs <- server.ListenAndServe()
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	m := <-errs
	log.Println("trigger", m)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("shutting down...")
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
