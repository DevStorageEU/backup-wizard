package device

import (
	"bwizard/internal/pkg/inspect"
)

type ProtectionStatus string
type Kind string

const (
	ProtectionStatusOk      ProtectionStatus = "ok"
	ProtectionStatusWarning ProtectionStatus = "warning"
	ProtectionStatusError   ProtectionStatus = "error"

	KindServer   Kind = "server"
	KindComputer Kind = "computer"
	KindMobile   Kind = "mobile"
)

type Inspection struct {
	Hostname        string
	CPU             string
	Ram             string
	Disks           []string
	OperatingSystem string
	Agent           string
}

// NewInspectionFromDeviceInspection returns a new domain inspection from inspect.DeviceInspection
func NewInspectionFromDeviceInspection(inspection *inspect.DeviceInspection) *Inspection {
	return &Inspection{
		Hostname:        inspection.Hostname,
		CPU:             inspection.CPU,
		Ram:             inspection.Ram,
		Disks:           inspection.Disks,
		Agent:           inspection.Agent,
		OperatingSystem: inspection.OperatingSystem,
	}
}
