package device

import (
	"bwizard/internal/pkg/wizard/domain/entity/device"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetDeviceByID(ctx context.Context, id uuid.UUID) (*device.Device, error)
	SaveDevice(ctx context.Context, device *device.Device) error
}
