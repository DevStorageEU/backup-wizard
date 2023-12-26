package wizardapi

import (
	"bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/application"
	deviceRepo "bwizard/internal/pkg/wizard/domain/repo/device"
	deviceValue "bwizard/internal/pkg/wizard/domain/valueobject/device"
	"bwizard/internal/pkg/wizardapi/mappings"
	"context"
)

type Handler struct {
	app              application.Application
	deviceRepository deviceRepo.Repository
}

var _ wizard.StrictServerInterface = &Handler{}

func NewHandler(app application.Application, deviceRepository deviceRepo.Repository) *Handler {
	return &Handler{
		app:              app,
		deviceRepository: deviceRepository,
	}
}

func (h *Handler) FindDevices(ctx context.Context, request wizard.FindDevicesRequestObject) (wizard.FindDevicesResponseObject, error) {
	devices, err := h.deviceRepository.GetDevices(ctx)
	if err != nil {
		// TODO: Add error handling here
		return nil, err
	}

	mappedDevices := mappings.MapDevices(devices)

	return wizard.FindDevices200JSONResponse{
		Success: true,
		Data:    mappedDevices,
	}, nil
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
