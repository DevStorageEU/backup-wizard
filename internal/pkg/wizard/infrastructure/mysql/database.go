package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"os"
)

func Connect(logger *zerolog.Logger) (*sql.DB, error) {
	logger.Info().Msg("try to connect to database...")

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
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
