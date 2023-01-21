package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markusryoti/werk/internal/auth"
	"github.com/markusryoti/werk/internal/postgres"
	"github.com/markusryoti/werk/internal/rest"
	"github.com/markusryoti/werk/internal/service"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger := initLogger()
	logger.Info("logger werkingss")

	authClient, err := auth.NewClient(ctx, logger)
	if err != nil {
		logger.Fatalw("Error instantiating auth client", "error", err)
	}

	repo, err := postgres.NewPostgresRepo(logger)
	if err != nil {
		logger.Fatalw("error instantiating postgres", "error", err)
	}
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	svc := service.NewService(repo)

	handler := rest.NewRouter(router, svc, logger, authClient)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler.Get(),
	}

	logger.Fatalf("error while listening: %s", s.ListenAndServe())
}

func initLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Panicf("Couldn't create a logger: %s", err.Error())
	}

	defer logger.Sync()
	sugar := logger.Sugar()

	return sugar
}
