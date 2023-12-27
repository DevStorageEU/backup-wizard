package device

import (
	"bwizard/internal/pkg/wizard/domain/valueobject/device"
	"github.com/gofrs/uuid/v5"
	"time"
)

type Device struct {
	ID              uuid.UUID
	Name            string
	Hostname        string
	Kind            device.Kind
	Protection      device.ProtectionStatus
	LastBackup      *time.Time
	CPU             string
	Ram             string
	Disks           []string
	IPs             []string
	OperatingSystem string
	Agent           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
