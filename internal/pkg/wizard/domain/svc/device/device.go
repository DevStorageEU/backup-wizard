package device

import (
	"bwizard/internal/pkg/wizard/domain/valueobject/device"
)

type InspectionService interface {
	Inspect(ips []device.IPAddress) (*device.Inspection, error)
}
