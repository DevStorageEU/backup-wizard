package repo

import (
	"bwizard/internal/pkg/wizard/domain/entity"
	"bwizard/internal/pkg/wizard/domain/repo"
	"context"
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const (
	DeviceDialect   = "mysql"
	DeviceTableName = "devices"
)

type DeviceRepoImpl struct {
	db *sql.DB
}

var _ repo.DeviceRepository = &DeviceRepoImpl{}

// NewDeviceRepository returns a new mysql device repository
func NewDeviceRepository(db *sql.DB) *DeviceRepoImpl {
	return &DeviceRepoImpl{
		db: db,
	}
}

// GetDeviceByID returns the device matching the given id from the database or nil
// if the device doesn't exist
func (i *DeviceRepoImpl) GetDeviceByID(ctx context.Context, id uuid.UUID) (*entity.Device, error) {
	query, _, err := goqu.
		Dialect(DeviceDialect).
		From(DeviceTableName).
		Select(&entity.Device{}).
		ToSQL()

	if err != nil {
		return nil, err
	}

	var device entity.Device
	row := i.db.QueryRowContext(ctx, query)

	switch rowErr := row.Scan(&device); {
	case errors.Is(rowErr, sql.ErrNoRows):
		return nil, nil
	case rowErr == nil:
		return &device, nil
	default:
		return nil, rowErr
	}
}

// SaveDevice inserts or updates a device in the database
func (i *DeviceRepoImpl) SaveDevice(ctx context.Context, device *entity.Device) error {
	insertSQL, _, _ := goqu.
		Dialect(DeviceDialect).
		Insert(DeviceTableName).
		Rows(device).
		OnConflict(goqu.DoUpdate("key", goqu.Record{"updated_at": goqu.L("NOW()")})).
		ToSQL()

	_, err := i.db.ExecContext(ctx, insertSQL)
	if err != nil {
		return err
	}

	return nil
}
