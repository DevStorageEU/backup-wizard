package device

import (
	"bwizard/internal/pkg/wizard/domain/valueobject/device"
)

type InspectionService interface {
	Inspect(ips []string) (*device.Inspection, error)
}
