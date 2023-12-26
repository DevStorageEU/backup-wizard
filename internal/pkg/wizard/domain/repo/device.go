package repo

import (
	"bwizard/internal/pkg/wizard/domain/entity"
	"context"
	"github.com/google/uuid"
)

type DeviceRepository interface {
	GetDeviceByID(ctx context.Context, id uuid.UUID) (*entity.Device, error)
	SaveDevice(ctx context.Context, device *entity.Device) error
}
