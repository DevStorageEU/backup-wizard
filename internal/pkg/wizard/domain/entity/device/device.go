package device

import (
	"bwizard/internal/pkg/wizard/domain/valueobject/device"
	"github.com/google/uuid"
	"time"
)

type Device struct {
	ID         uuid.UUID               `db:"id"`
	Name       string                  `db:"name"`
	Kind       device.Kind             `db:"kind"`
	Protection device.ProtectionStatus `db:"protection"`
	LastBackup *time.Time              `db:"last_backup" goqu:"omitempty"`
	IPs        []device.IPAddress      `db:"ips"`
	Agent      string                  `db:"agent"`
	CreatedAt  time.Time               `db:"created_at"`
	UpdatedAt  time.Time               `db:"updated_at"`
}
