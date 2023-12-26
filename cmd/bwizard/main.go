package main

import (
	"bwizard/internal/pkg/wizard/application"
	"bwizard/internal/pkg/wizard/infrastructure/mysql"
	"bwizard/internal/pkg/wizardapi"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	db, err := mysql.Connect(&logger)
	if err != nil {
		logger.Error().Msg(err.Error())
		os.Exit(1)
	}

	app := application.NewApplication(db)

	wizardapi.StartAPIServer(app)
}
