package main

import (
	"bwizard/internal/pkg/ssh"
	"bwizard/internal/pkg/wizard/application"
	"bwizard/internal/pkg/wizard/infrastructure/mysql"
	"bwizard/internal/pkg/wizardapi"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	ApplicationDir = "/var/lib/bwizard"
)

func main() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	if err := os.Mkdir(ApplicationDir, 0755); err != nil {
		if !os.IsExist(err) {
			logger.Error().Msg(err.Error())
			os.Exit(1)
		}
	}

	if err := ssh.CreateSSHKeyPairIfNotExists(ApplicationDir); err != nil {
		logger.Error().Msg(err.Error())
		os.Exit(1)
	}

	db, err := mysql.Connect(&logger)
	if err != nil {
		logger.Error().Msg(err.Error())
		os.Exit(1)
	}

	app := application.NewApplication(db)

	wizardapi.StartAPIServer(app)
}
