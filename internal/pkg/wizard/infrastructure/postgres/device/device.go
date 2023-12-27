package device

import (
	"bwizard/internal/pkg/wizard/domain/device"
	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
	"time"
)

type Device struct {
	ID              uuid.UUID               `db:"id"`
	Name            string                  `db:"name"`
	Hostname        string                  `db:"hostname"`
	Kind            device.Kind             `db:"kind"`
	Protection      device.ProtectionStatus `db:"protection"`
	LastBackup      *time.Time              `db:"last_backup" goqu:"omitempty"`
	CPU             string                  `db:"cpu"`
	Ram             string                  `db:"ram"`
	Disks           pq.StringArray          `db:"disks"`
	IPs             pq.StringArray          `db:"ips"`
	OperatingSystem string                  `db:"operating_system"`
	Agent           string                  `db:"agent"`
	CreatedAt       time.Time               `db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt       time.Time               `db:"updated_at" goqu:"skipinsert"`
}

func MapModel(deviceEntity *Device) *device.Device {
	return &device.Device{
		ID:              deviceEntity.ID,
		Name:            deviceEntity.Name,
		Hostname:        deviceEntity.Hostname,
		Kind:            deviceEntity.Kind,
		Protection:      deviceEntity.Protection,
		LastBackup:      deviceEntity.LastBackup,
		CPU:             deviceEntity.CPU,
		Ram:             deviceEntity.Ram,
		Disks:           deviceEntity.Disks,
		IPs:             deviceEntity.IPs,
		OperatingSystem: deviceEntity.OperatingSystem,
		Agent:           deviceEntity.Agent,
		CreatedAt:       deviceEntity.CreatedAt,
		UpdatedAt:       deviceEntity.UpdatedAt,
	}
}

func MapDevice(deviceEntity *device.Device) *Device {
	return &Device{
		ID:              deviceEntity.ID,
		Name:            deviceEntity.Name,
		Hostname:        deviceEntity.Hostname,
		Kind:            deviceEntity.Kind,
		Protection:      deviceEntity.Protection,
		LastBackup:      deviceEntity.LastBackup,
		CPU:             deviceEntity.CPU,
		Ram:             deviceEntity.Ram,
		Disks:           deviceEntity.Disks,
		IPs:             deviceEntity.IPs,
		OperatingSystem: deviceEntity.OperatingSystem,
		Agent:           deviceEntity.Agent,
		CreatedAt:       deviceEntity.CreatedAt,
		UpdatedAt:       deviceEntity.UpdatedAt,
	}
}
