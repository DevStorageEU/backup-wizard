package mappings

import (
	"bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/domain/entity/device"
)

func MapDevice(device *device.Device) wizard.Device {
	ips := make([]string, 0, len(device.IPs))
	for _, ip := range ips {
		ips = append(ips, ip)
	}

	return wizard.Device{
		Id:         device.ID,
		Name:       device.Name,
		Kind:       wizard.DeviceKind(device.Kind),
		Protection: wizard.ProtectionStatus(device.Protection),
		Agent:      device.Agent,
		Ips:        ips,
		LastBackup: device.LastBackup,
	}
}

func MapDevices(devices []*device.Device) []wizard.Device {
	mappedDevices := make([]wizard.Device, 0, len(devices))
	for _, deviceEntity := range devices {
		mappedDevices = append(mappedDevices, MapDevice(deviceEntity))
	}

	return mappedDevices
}
