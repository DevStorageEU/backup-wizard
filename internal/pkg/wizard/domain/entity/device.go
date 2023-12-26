package entity

import (
	"github.com/google/uuid"
	"time"
)

type ProtectionStatus string
type DeviceKind string

const (
	ProtectionStatusOk      ProtectionStatus = "ok"
	ProtectionStatusWarning ProtectionStatus = "warning"
	ProtectionStatusError   ProtectionStatus = "error"

	DeviceKindServer   DeviceKind = "server"
	DeviceKindComputer DeviceKind = "computer"
	DeviceKindMobile   DeviceKind = "mobile"
)

type OperatingSystem struct {
	Name           string `db:"name"`
	CPU            string `db:"cpu"`
	Ram            string `db:"ram"`
	TotalHardDrive string `db:"ram"`
}

type Device struct {
	ID         uuid.UUID        `db:"id"`
	Name       string           `db:"name"`
	Kind       DeviceKind       `db:"kind"`
	Protection ProtectionStatus `db:"protection"`
	LastBackup *time.Time       `db:"last_backup" goqu:"omitempty"`
	IPs        []string         `db:"ips"`
	Agent      string           `db:"agent"`
	CreatedAt  time.Time        `db:"created_at"`
	UpdatedAt  time.Time        `db:"updated_at"`
}
