package wizardapi

import (
	"bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/application"
	deviceRepo "bwizard/internal/pkg/wizard/domain/device"
	"bwizard/internal/pkg/wizardapi/mappings"
	"context"
	"github.com/rs/zerolog"
	"os"
)

type Handler struct {
	logger           *zerolog.Logger
	app              application.Application
	deviceRepository deviceRepo.Repository
}

var _ wizard.StrictServerInterface = &Handler{}

func NewHandler(logger *zerolog.Logger, app application.Application, deviceRepository deviceRepo.Repository) *Handler {
	return &Handler{
		logger:           logger,
		app:              app,
		deviceRepository: deviceRepository,
	}
}

func (h *Handler) FindSSHKeys(ctx context.Context, request wizard.FindSSHKeysRequestObject) (wizard.FindSSHKeysResponseObject, error) {
	// TODO: Replace with some kind of ssh keys repository

	sshKey, err := os.ReadFile("/var/lib/bwizard/id_ed25519.pub")
	if os.IsNotExist(err) {
		return wizard.FindSSHKeys200JSONResponse{
			Success: true,
			Data:    make([]wizard.SSHKey, 0),
		}, nil
	}

	if err != nil {
		// TODO: Add error handling here
		return nil, err
	}

	return wizard.FindSSHKeys200JSONResponse{
		Success: true,
		Data:    []wizard.SSHKey{string(sshKey)},
	}, nil
}

func (h *Handler) FindDevices(ctx context.Context, request wizard.FindDevicesRequestObject) (wizard.FindDevicesResponseObject, error) {
	devices, err := h.deviceRepository.GetDevices(ctx)
	if err != nil {
		// TODO: Add error handling here
		return nil, err
	}

	mappedDevices, err := mappings.MapDevices(devices)
	if err != nil {
		h.logger.Error().Msgf("%s", err.Error())

		// TODO: Add error handling here
		return nil, err
	}

	return wizard.FindDevices200JSONResponse{
		Success: true,
		Data:    mappedDevices,
	}, nil
}

func (h *Handler) RegisterDevice(ctx context.Context, request wizard.RegisterDeviceRequestObject) (wizard.RegisterDeviceResponseObject, error) {
	device, err := h.app.SetupDevice(ctx, request.Body.Ips, deviceRepo.Kind(request.Body.Kind), request.Body.Name)
	if err != nil {
		h.logger.Error().Msgf("%s", err.Error())

		// TODO: Add error handling here
		return nil, err
	}

	mappedDevice, err := mappings.MapDevice(device)
	if err != nil {
		h.logger.Error().Msgf("%s", err.Error())

		// TODO: Add error handling here
		return nil, err
	}

	return wizard.RegisterDevice200JSONResponse{
		Success: true,
		Data:    mappedDevice,
	}, nil
}

func (h *Handler) GetDevicesId(ctx context.Context, request wizard.GetDevicesIdRequestObject) (wizard.GetDevicesIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
