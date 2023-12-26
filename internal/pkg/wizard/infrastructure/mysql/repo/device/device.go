package device

import (
	"bwizard/internal/pkg/wizard/domain/entity/device"
	device2 "bwizard/internal/pkg/wizard/domain/repo/device"
	"context"
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

const (
	Dialect   = "mysql"
	TableName = "devices"
)

type Impl struct {
	db *sql.DB
}

var _ device2.Repository = &Impl{}

// NewRepository returns a new mysql device repository
func NewRepository(db *sql.DB) *Impl {
	return &Impl{
		db: db,
	}
}

// GetDeviceByID returns the device matching the given id from the database or nil
// if the device doesn't exist
func (i *Impl) GetDeviceByID(ctx context.Context, id uuid.UUID) (*device.Device, error) {
	query, _, err := goqu.
		Dialect(Dialect).
		From(TableName).
		Select(&device.Device{}).
		ToSQL()

	if err != nil {
		return nil, err
	}

	var deviceEntity device.Device
	row := i.db.QueryRowContext(ctx, query)

	switch rowErr := row.Scan(&deviceEntity); {
	case errors.Is(rowErr, sql.ErrNoRows):
		return nil, nil
	case rowErr == nil:
		return &deviceEntity, nil
	default:
		return nil, rowErr
	}
}

// SaveDevice inserts or updates a device in the database
func (i *Impl) SaveDevice(ctx context.Context, device *device.Device) error {
	insertSQL, _, _ := goqu.
		Dialect(Dialect).
		Insert(TableName).
		Rows(device).
		OnConflict(goqu.DoUpdate("key", goqu.Record{"updated_at": goqu.L("NOW()")})).
		ToSQL()

	_, err := i.db.ExecContext(ctx, insertSQL)
	if err != nil {
		return err
	}

	return nil
}
