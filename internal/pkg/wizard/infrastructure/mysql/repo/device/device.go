package device

import (
	"bwizard/internal/pkg/wizard/domain/entity/device"
	device2 "bwizard/internal/pkg/wizard/domain/repo/device"
	deviceModel "bwizard/internal/pkg/wizard/infrastructure/mysql/model/device"
	"context"
	"database/sql"
	"errors"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

const (
	Dialect   = "postgres"
	TableName = "devices"
)

type Impl struct {
	logger *zerolog.Logger
	db     *sqlx.DB
}

var _ device2.Repository = &Impl{}

// NewRepository returns a new mysql device repository
func NewRepository(logger *zerolog.Logger, db *sqlx.DB) *Impl {
	return &Impl{
		logger: logger,
		db:     db,
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

// GetDevices returns all registered devices
func (i *Impl) GetDevices(ctx context.Context) ([]*device.Device, error) {
	query, _, err := goqu.
		Dialect(Dialect).
		From(TableName).
		Select(&deviceModel.Device{}).
		ToSQL()

	if err != nil {
		return nil, err
	}

	rows, err := i.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	devices := make([]*device.Device, 0)

	for rows.Next() {
		var model deviceModel.Device
		if scanErr := rows.StructScan(&model); scanErr != nil {
			return nil, scanErr
		}

		deviceEntity := deviceModel.MapModel(&model)
		devices = append(devices, deviceEntity)
	}

	return devices, nil
}

// SaveDevice inserts or updates a device in the database
func (i *Impl) SaveDevice(ctx context.Context, device *device.Device) error {
	i.logger.Debug().Msgf("try to save device %s", device.ID.String())

	model := deviceModel.MapDevice(device)

	insertSQL, _, _ := goqu.
		Dialect(Dialect).
		Insert(TableName).
		Rows(model).
		OnConflict(goqu.DoUpdate("id", goqu.Record{"updated_at": goqu.L("NOW()")})).
		ToSQL()

	i.logger.Debug().Msgf("got insert query %s", insertSQL)

	_, err := i.db.ExecContext(ctx, insertSQL)
	if err != nil {
		return err
	}

	return nil
}
