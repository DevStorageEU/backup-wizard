package device

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Builder struct {
	*Device
}

func NewBuilder() *Builder {
	return &Builder{
		Device: &Device{},
	}
}

func (b *Builder) ID(v uuid.UUID) *Builder {
	b.Device.ID = v

	return b
}

func (b *Builder) Name(v string) *Builder {
	b.Device.Name = v

	return b
}

func (b *Builder) Hostname(v string) *Builder {
	b.Device.Hostname = v

	return b
}

func (b *Builder) Kind(v Kind) *Builder {
	b.Device.Kind = v

	return b
}

func (b *Builder) Protection(v ProtectionStatus) *Builder {
	b.Device.Protection = v

	return b
}

func (b *Builder) LastBackup(v *time.Time) *Builder {
	b.Device.LastBackup = v

	return b
}

func (b *Builder) CPU(v string) *Builder {
	b.Device.CPU = v

	return b
}

func (b *Builder) Ram(v string) *Builder {
	b.Device.Ram = v

	return b
}

func (b *Builder) Disks(v []string) *Builder {
	b.Device.Disks = v

	return b
}

func (b *Builder) IPs(v []string) *Builder {
	b.Device.IPs = v

	return b
}

func (b *Builder) OperatingSystem(v string) *Builder {
	b.Device.OperatingSystem = v

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

func (b *Builder) Build() (*Device, error) {
	if b.Device.ID.IsNil() {
		deviceID, err := uuid.NewV7()
		if err != nil {
			return nil, err
		}

		b.Device.ID = deviceID
	}

	if b.Device.CreatedAt.IsZero() {
		b.Device.CreatedAt = time.Now()
	}

	if b.Device.UpdatedAt.IsZero() {
		b.Device.UpdatedAt = time.Now()
	}

	return &Device{
		ID:              b.Device.ID,
		Name:            b.Device.Name,
		Hostname:        b.Device.Hostname,
		Kind:            b.Device.Kind,
		Protection:      b.Device.Protection,
		LastBackup:      b.Device.LastBackup,
		CPU:             b.Device.CPU,
		Ram:             b.Device.Ram,
		Disks:           b.Device.Disks,
		IPs:             b.Device.IPs,
		Agent:           b.Device.Agent,
		OperatingSystem: b.Device.OperatingSystem,
		CreatedAt:       b.Device.CreatedAt,
		UpdatedAt:       b.Device.UpdatedAt,
	}, nil
}
