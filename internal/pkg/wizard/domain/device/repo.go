package device

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetDeviceByID(ctx context.Context, id uuid.UUID) (*Device, error)
	GetDevices(ctx context.Context) ([]*Device, error)
	SaveDevice(ctx context.Context, device *Device) error
}
