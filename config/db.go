package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}

	return db
}
