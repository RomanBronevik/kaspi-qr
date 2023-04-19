package old

import (
	"context"
	"errors"
	"kaspi-qr/internal/adapters/db"
	_ "kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) CreateDevice(ctx context.Context, device *entities.CreateDeviceDTO) error {
	q := `
		INSERT INTO device (device_id, token, organization_bin) 
		VALUES ($1, $2, $3)`

	return d.db.Exec(ctx, q, device.DeviceId, device.Token, device.OrganizationBin)
}

func (d *St) FindAllDevices(ctx context.Context) ([]*entities.Device, error) {
	q := `
		SELECT device_id, token,  organization_bin
		FROM device`

	rows, err := d.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := make([]*entities.Device, 0)

	for rows.Next() {
		dev := &entities.Device{}

		err = rows.Scan(&dev.DeviceId, &dev.Token, &dev.OrganizationBin)
		if err != nil {
			return nil, err
		}

		devices = append(devices, dev)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

func (d *St) FindOneDevice(ctx context.Context, token string) (*entities.Device, error) {
	q := `
		SELECT device_id, token, organization_bin
		FROM device
		WHERE token = $1`

	dev := &entities.Device{}

	err := d.db.QueryRow(ctx, q, token).Scan(&dev.DeviceId, &dev.Token, &dev.OrganizationBin)
	if err != nil {
		if errors.Is(err, db.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return dev, nil
}

func (d *St) DeleteDevice(ctx context.Context, bin string, token string) error {
	q := `
		DELETE FROM device
		WHERE organization_bin = $1 AND token = $2;`

	return d.db.Exec(ctx, q, bin, token)
}
