package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
)

func Connect(logger *zerolog.Logger) (*sqlx.DB, error) {
	logger.Info().Msg("try to connect to database...")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"))

	logger.Debug().Msgf("dsn: %s", dsn)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	logger.Info().Msg("successfully connected to database")

	return db, nil
}
