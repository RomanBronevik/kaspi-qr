package old

import (
	"context"
	"errors"
	_ "kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) CreateDevice(ctx context.Context, device *entities.CreateDeviceDTO) error {
	q := `
		INSERT INTO device (device_id, token, organization_bin) 
		VALUES ($1, $2, $3)`

	return d.DbExec(ctx, q, device.DeviceId, device.Token, device.OrganizationBin)
}

func (d *St) FindAllDevices(ctx context.Context) ([]*entities.Device, error) {
	q := `
		SELECT device_id, token,  organization_bin
		FROM device`

	rows, err := d.DbQuery(ctx, q)
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

	err := d.DbQueryRow(ctx, q, token).Scan(&dev.DeviceId, &dev.Token, &dev.OrganizationBin)
	if err != nil {
		if errors.Is(err, dopErrs.NoRows) {
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

	return d.DbExec(ctx, q, bin, token)
}
