package wizardapi

import (
	api "bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/application"
	"flag"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net"
)

const (
	DefaultPort = "8099"
)

func StartAPIServer(app application.Application) {
	port := flag.String("port", DefaultPort, "Default port to start the api and ui on")
	flag.Parse()

	handler := NewHandler(app)

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	//e.Use(middleware.OapiRequestValidator())

	strictAPIHandler := api.NewStrictHandler(handler, nil)
	api.RegisterHandlers(e, strictAPIHandler)

	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", *port)))
}
