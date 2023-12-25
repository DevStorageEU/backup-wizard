package wizardapi

import (
	"bwizard/api/openapi/wizard"
	"github.com/labstack/echo/v4"
	openapiTypes "github.com/oapi-codegen/runtime/types"
)

type Handler struct{}

var _ wizard.ServerInterface = &Handler{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetDevices(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PostDevices(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetDevicesId(ctx echo.Context, id openapiTypes.UUID) error {
	//TODO implement me
	panic("implement me")
}
