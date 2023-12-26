package device

import (
	"bwizard/internal/pkg/wizard/domain/valueobject/device"
	"github.com/google/uuid"
	"time"
)

type Builder struct {
	*Device
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) ID(v uuid.UUID) *Builder {
	b.Device.ID = v

	return b
}

func (b *Builder) Name(v string) *Builder {
	b.Device.Name = v

	return b
}

func (b *Builder) Kind(v device.Kind) *Builder {
	b.Device.Kind = v

	return b
}

func (b *Builder) Protection(v device.ProtectionStatus) *Builder {
	b.Device.Protection = v

	return b
}

func (b *Builder) LastBackup(v *time.Time) *Builder {
	b.Device.LastBackup = v

	return b
}

func (b *Builder) IPs(v []device.IPAddress) *Builder {
	b.Device.IPs = v

	return b
}

func (b *Builder) Agent(v string) *Builder {
	b.Device.Agent = v

	return b
}

func (b *Builder) CreatedAt(v time.Time) *Builder {
	b.Device.CreatedAt = v

	return b
}

func (b *Builder) UpdatedAt(v time.Time) *Builder {
	b.Device.UpdatedAt = v

	return b
}

func (b *Builder) Build() *Device {
	CreatedAt := b.Device.CreatedAt
	UpdatedAt := b.Device.UpdatedAt

	if CreatedAt.IsZero() {
		b.Device.CreatedAt = time.Now()
	}

	if UpdatedAt.IsZero() {
		b.Device.UpdatedAt = time.Now()
	}

	return &Device{
		ID:         b.Device.ID,
		Name:       b.Device.Name,
		Kind:       b.Device.Kind,
		Protection: b.Device.Protection,
		LastBackup: b.Device.LastBackup,
		IPs:        b.Device.IPs,
		Agent:      b.Device.Agent,
		CreatedAt:  b.Device.CreatedAt,
		UpdatedAt:  b.Device.UpdatedAt,
	}
}
