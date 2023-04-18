package pg

import (
	"context"
	"errors"

	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (d *St) CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error {
	q := `
		INSERT INTO organization (name, bin) 
		VALUES ($1, $2)`

	return d.db.Exec(ctx, q, organization.Name, organization.Bin)
}

func (d *St) FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error) {
	q := `
		SELECT name, bin FROM organization`

	rows, err := d.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	organizations := make([]entities.Organization, 0)

	for rows.Next() {
		var org entities.Organization

		err = rows.Scan(&org.Name, &org.Bin)
		if err != nil {
			return nil, d.db.HErr(err)
		}

		organizations = append(organizations, org)
	}
	if err = rows.Err(); err != nil {
		return nil, d.db.HErr(err)
	}

	return organizations, nil
}

func (d *St) FindOneOrganization(ctx context.Context, bin string) (entities.Organization, error) {
	q := `
		SELECT NAME, BIN FROM organization WHERE BIN = &1`

	var org entities.Organization

	err := d.db.QueryRow(ctx, q, bin).Scan(&org.Name, &org.Bin)
	if err != nil {
		err = d.db.HErr(err)
		if !errors.Is(err, db.ErrNoRows) {
			return entities.Organization{}, err
		}
	}

	return org, nil
}

// func (r *St) UpdateOrganization(ctx context.Context, organization entities.Organization) error {
//	//TODO implement me
//	panic("implement me")
// }

func (d *St) DeleteOrganization(ctx context.Context, bin string) error {
	q := `
		DELETE FROM organization
		WHERE bin = $1;`

	if err := d.db.Exec(ctx, q, bin); err != nil {
		return err
	}

	return nil
}
