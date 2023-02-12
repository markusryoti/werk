package postgres

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type PostgresRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewPostgresRepo(logger *zap.SugaredLogger) (*PostgresRepo, error) {
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", host, user, dbName, password, port))
	if err != nil {
		return nil, err
	}

	logger.Info("Postgres connected")

	return &PostgresRepo{
		db:     db,
		logger: logger,
	}, nil
}
