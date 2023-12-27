package mappings

import (
	"bwizard/api/openapi/wizard"
	"bwizard/internal/pkg/wizard/domain/entity/device"
	"github.com/google/uuid"
)

func MapDevice(device *device.Device) (wizard.Device, error) {
	ips := make([]string, 0, len(device.IPs))
	for _, ip := range device.IPs {
		ips = append(ips, string(ip))
	}

	disks := make([]string, 0, len(device.Disks))
	for _, disk := range device.Disks {
		disks = append(disks, string(disk))
	}

	id, err := uuid.FromBytes(device.ID.Bytes())
	if err != nil {
		return wizard.Device{}, err
	}

	return wizard.Device{
		Id:              id,
		Name:            device.Name,
		Hostname:        device.Hostname,
		Kind:            wizard.DeviceKind(device.Kind),
		Protection:      wizard.ProtectionStatus(device.Protection),
		Cpu:             device.CPU,
		Ram:             device.Ram,
		Disks:           disks,
		Ips:             ips,
		LastBackup:      device.LastBackup,
		Agent:           device.Agent,
		OperatingSystem: device.OperatingSystem,
	}, nil
}

func MapDevices(devices []*device.Device) ([]wizard.Device, error) {
	mappedDevices := make([]wizard.Device, 0, len(devices))
	for _, deviceEntity := range devices {
		mappedDevice, err := MapDevice(deviceEntity)
		if err != nil {
			return nil, err
		}

		mappedDevices = append(mappedDevices, mappedDevice)
	}

	return mappedDevices, nil
}
