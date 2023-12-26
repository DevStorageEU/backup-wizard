package application

import (
	deviceSvcImpl "bwizard/internal/pkg/wizard/application/svc/device"
	"bwizard/internal/pkg/wizard/domain/entity/device"
	deviceRepo "bwizard/internal/pkg/wizard/domain/repo/device"
	deviceSvc "bwizard/internal/pkg/wizard/domain/svc/device"
	deviceValue "bwizard/internal/pkg/wizard/domain/valueobject/device"
	deviceRepoImpl "bwizard/internal/pkg/wizard/infrastructure/mysql/repo/device"
	"database/sql"
)

type Application interface {
	// SetupDevice setups a new device
	SetupDevice(ips []deviceValue.IPAddress, kind deviceValue.Kind, name *string) (*device.Device, error)
}

type Impl struct {
	deviceRepo          deviceRepo.Repository
	deviceInspectionSvc deviceSvc.InspectionService
}

var _ Application = &Impl{}

// NewApplication returns a new application
func NewApplication(db *sql.DB) *Impl {
	deviceRepository := deviceRepoImpl.NewRepository(db)
	deviceInspectionSvc := deviceSvcImpl.NewInspectionSvc()

	return &Impl{
		deviceRepo:          deviceRepository,
		deviceInspectionSvc: deviceInspectionSvc,
	}
}

// SetupDevice setups a new device
func (i *Impl) SetupDevice(ips []deviceValue.IPAddress, kind deviceValue.Kind, name *string) (*device.Device, error) {
	inspection, err := i.deviceInspectionSvc.Inspect(ips)
	if err != nil {
		return nil, err
	}

	if inspection == nil {
		// TODO: Throw some kind of error if inspection failed
	}

	if name == nil {
		name = &inspection.Hostname
	}

	deviceEntity := device.NewBuilder().
		Name(*name).
		Kind(kind).
		IPs(ips).
		Build()

	return deviceEntity, nil
}
