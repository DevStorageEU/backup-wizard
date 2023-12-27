package application

import (
	deviceSvcImpl "bwizard/internal/pkg/wizard/application/svc/device"
	"bwizard/internal/pkg/wizard/domain/entity/device"
	deviceRepo "bwizard/internal/pkg/wizard/domain/repo/device"
	deviceSvc "bwizard/internal/pkg/wizard/domain/svc/device"
	deviceValue "bwizard/internal/pkg/wizard/domain/valueobject/device"
	deviceRepoImpl "bwizard/internal/pkg/wizard/infrastructure/mysql/repo/device"
	"context"
	"errors"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Application interface {
	// SetupDevice setups a new device
	SetupDevice(ctx context.Context, ips []string, kind deviceValue.Kind, name *string) (*device.Device, error)
}

type Impl struct {
	logger              *zerolog.Logger
	DeviceRepo          deviceRepo.Repository
	DeviceInspectionSvc deviceSvc.InspectionService
}

var _ Application = &Impl{}

// NewApplication returns a new application
func NewApplication(logger *zerolog.Logger, db *sqlx.DB) *Impl {
	deviceRepository := deviceRepoImpl.NewRepository(logger, db)
	deviceInspectionSvc := deviceSvcImpl.NewInspectionSvc(logger)

	return &Impl{
		logger:              logger,
		DeviceRepo:          deviceRepository,
		DeviceInspectionSvc: deviceInspectionSvc,
	}
}

// SetupDevice setups a new device
func (i *Impl) SetupDevice(ctx context.Context, ips []string, kind deviceValue.Kind, name *string) (*device.Device, error) {
	inspection, err := i.DeviceInspectionSvc.Inspect(ips)
	if err != nil {
		return nil, err
	}

	i.logger.Debug().Msgf("got inspection %+v", inspection)

	if inspection == nil {
		// TODO: Throw some kind of error if inspection failed
		return nil, errors.New("inspection missing")
	}

	if name == nil {
		name = &inspection.Hostname
	}

	deviceEntity, err := device.NewBuilder().
		Name(*name).
		Hostname(inspection.Hostname).
		Kind(kind).
		Protection(deviceValue.ProtectionStatusOk).
		CPU(inspection.CPU).
		Ram(inspection.Ram).
		Disks(inspection.Disks).
		Agent(inspection.Agent).
		OperatingSystem(inspection.OperatingSystem).
		IPs(ips).
		Build()

	if err != nil {
		return nil, err
	}

	i.logger.Debug().Msgf("got device entity %+v", deviceEntity)

	if saveErr := i.DeviceRepo.SaveDevice(ctx, deviceEntity); saveErr != nil {
		return nil, saveErr
	}

	return deviceEntity, nil
}
