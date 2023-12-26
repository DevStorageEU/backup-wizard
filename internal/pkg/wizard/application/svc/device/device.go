package device

import (
	"bwizard/internal/pkg/inspect"
	"bwizard/internal/pkg/ssh"
	deviceSvc "bwizard/internal/pkg/wizard/domain/svc/device"
	deviceValue "bwizard/internal/pkg/wizard/domain/valueobject/device"
)

type InspectionSvcImpl struct {
}

var _ deviceSvc.InspectionService = &InspectionSvcImpl{}

// Inspect gains information from a backup device such like the operating system or the agent
func (i *InspectionSvcImpl) Inspect(ips []deviceValue.IPAddress) (*deviceValue.Inspection, error) {
	credentials := ssh.Credentials{
		Username: "root",
	}

	var clients []*ssh.Client
	for _, ip := range ips {

		client, err := ssh.NewClient(string(ip), 22, &credentials)
		if err != nil {
			return nil, err
		}

		if err := client.Connect(); err != nil {
			continue
		}

		clients = append(clients, client)

		inspection, err := inspect.Device(client)
		if err != nil {
			return nil, err
		}

		return deviceValue.NewInspectionFromDeviceInspection(inspection), nil
	}

	defer func(clients []*ssh.Client) {
		for _, client := range clients {
			err := client.Close()
			if err != nil {
				panic(err)
			}
		}
	}(clients)

	return nil, nil
}
