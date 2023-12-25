package wizardapi

import (
	api "bwizard/api/openapi/wizard"
	"flag"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net"
)

const (
	DefaultPort = "8099"
)

func StartAPIServer() {
	port := flag.String("port", DefaultPort, "Default port to start the api and ui on")
	flag.Parse()

	handler := NewHandler()

	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	//e.Use(middleware.OapiRequestValidator())

	api.RegisterHandlers(e, handler)

	e.Logger.Fatal(e.Start(net.JoinHostPort("127.0.0.1", *port)))
}
