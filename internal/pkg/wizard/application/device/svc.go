package device

import (
	"bwizard/internal/pkg/inspect"
	"bwizard/internal/pkg/ssh"
	"bwizard/internal/pkg/wizard/domain/device"
	"github.com/rs/zerolog"
)

const (
	privateKeyPath = "/var/lib/bwizard/id_ed25519"
)

type InspectionSvcImpl struct {
	logger *zerolog.Logger
}

var _ device.InspectionService = &InspectionSvcImpl{}

func NewInspectionSvc(logger *zerolog.Logger) *InspectionSvcImpl {
	return &InspectionSvcImpl{
		logger: logger,
	}
}

// Inspect gains information from a backup device such like the operating system or the agent
func (i *InspectionSvcImpl) Inspect(ips []string) (*device.Inspection, error) {
	credentials := ssh.Credentials{
		Username:       "root",
		PrivateKeyPath: privateKeyPath,
	}

	i.logger.Debug().Msgf("try %d hosts", len(ips))

	var clients []*ssh.Client
	for _, ip := range ips {

		i.logger.Debug().Msgf("try host %s", ip)

		client, err := ssh.NewClient(ip, 22, &credentials)
		if err != nil {
			return nil, err
		}

		if connectErr := client.Connect(); connectErr != nil {
			i.logger.Debug().Msgf("skip connection to %s due to error: %s", ip, connectErr.Error())
			continue
		}

		i.logger.Debug().Msgf("successfully connected to %s", ip)

		clients = append(clients, client)

		inspection, err := inspect.Device(client)
		if err != nil {
			return nil, err
		}

		return device.NewInspectionFromDeviceInspection(inspection), nil
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
