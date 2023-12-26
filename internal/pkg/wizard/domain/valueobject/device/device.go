package device

import "bwizard/internal/pkg/inspect"

type ProtectionStatus string
type Kind string
type IPAddress string

const (
	ProtectionStatusOk      ProtectionStatus = "ok"
	ProtectionStatusWarning ProtectionStatus = "warning"
	ProtectionStatusError   ProtectionStatus = "error"

	KindServer   Kind = "server"
	KindComputer Kind = "computer"
	KindMobile   Kind = "mobile"
)

type OperatingSystem struct {
	Name           string `db:"name"`
	CPU            string `db:"cpu"`
	Ram            string `db:"ram"`
	TotalHardDrive string `db:"ram"`
}

type Inspection struct {
	Hostname string `db:"hostname"`
	CPU      string `db:"cpu"`
	Ram      string `db:"ram"`
}

// NewInspectionFromDeviceInspection returns a new domain inspection from inspect.DeviceInspection
func NewInspectionFromDeviceInspection(inspection *inspect.DeviceInspection) *Inspection {
	return &Inspection{
		Hostname: inspection.Hostname,
		CPU:      inspection.CPU,
		Ram:      inspection.Ram,
	}
}
