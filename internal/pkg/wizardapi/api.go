package wizardapi

import (
	api "bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/application"
	deviceRepo "bwizard/internal/pkg/wizard/domain/device"
	"flag"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"net"
)

const (
	DefaultPort = "8099"
)

func StartAPIServer(logger *zerolog.Logger, app application.Application, deviceRepository deviceRepo.Repository) {
	port := flag.String("port", DefaultPort, "Default port to start the api and ui on")
	flag.Parse()

	handler := NewHandler(logger, app, deviceRepository)

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	//e.Use(middleware.OapiRequestValidator())

	g := e.Group("/v1")

	strictAPIHandler := api.NewStrictHandler(handler, nil)
	api.RegisterHandlers(g, strictAPIHandler)

	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
