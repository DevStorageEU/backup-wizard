package wizardapi

import (
	"bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/application"
	deviceValue "bwizard/internal/pkg/wizard/domain/valueobject/device"
	"bwizard/internal/pkg/wizardapi/mappings"
	"context"
)

type Handler struct {
	app application.Application
}

var _ wizard.StrictServerInterface = &Handler{}

func NewHandler(app application.Application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) FindDevices(ctx context.Context, request wizard.FindDevicesRequestObject) (wizard.FindDevicesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) RegisterDevice(ctx context.Context, request wizard.RegisterDeviceRequestObject) (wizard.RegisterDeviceResponseObject, error) {
	ips := make([]deviceValue.IPAddress, 0)
	for _, ip := range ips {
		ips = append(ips, ip)
	}

	device, err := h.app.SetupDevice(ips, deviceValue.Kind(request.Body.Kind), request.Body.Name)
	if err != nil {
		// TODO: Add error handling here
		return nil, err
	}

	mappedDevice := mappings.MapDevice(device)

	return wizard.RegisterDevice200JSONResponse{
		Success: true,
		Data:    mappedDevice,
	}, nil
}

func (h *Handler) GetDevicesId(ctx context.Context, request wizard.GetDevicesIdRequestObject) (wizard.GetDevicesIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
