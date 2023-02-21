package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	_ "kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateDevice(ctx context.Context, device *entities.CreateDeviceDTO) error {
	q := `
		INSERT INTO device (device_id, token, organization_bin) 
		VALUES ($1, $2, $3) 
		RETURNING id`

	var id string

	if err := r.client.QueryRow(ctx, q, device.DeviceId, device.Token, device.OrganizationBin).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *St) FindAllDevices(ctx context.Context) (u []entities.Device, err error) {
	q := `
		SELECT id, device_id, token,  organization_bin FROM device`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	devices := make([]entities.Device, 0)

	for rows.Next() {
		var dev entities.Device

		err := rows.Scan(&dev.ID, &dev.DeviceId, &dev.Token, &dev.OrganizationBin)
		if err != nil {
			return nil, err
		}

		devices = append(devices, dev)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *St) FindOneDevice(ctx context.Context, OrganizationBin string) (entities.Device, error) {
	q := `
		SELECT id, device_id, token, organization_bin FROM public.device WHERE organization_bin = $1`

	//Trace

	var dev entities.Device
	err := r.client.QueryRow(ctx, q, OrganizationBin).Scan(&dev.ID, &dev.DeviceId, &dev.Token, &dev.OrganizationBin)
	if err != nil {
		return entities.Device{}, err
	}

	return dev, nil

}

// func (r *St) UpdateDevice(ctx context.Context, token string) error {
//
//		panic("implement me")
//	}
func (r *St) DeleteDevice(ctx context.Context, bin string, token string) error {
	q := `
		DELETE FROM device
		WHERE organization_bin = $1 AND token = $2;`

	var id string

	if err := r.client.QueryRow(ctx, q, bin, token).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}
