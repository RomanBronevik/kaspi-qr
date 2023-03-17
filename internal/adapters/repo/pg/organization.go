package pg

import (
	"context"
	"errors"

	"kaspi-qr/internal/adapters/db"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error {
	q := `
		INSERT INTO organization (name, bin) 
		VALUES ($1, $2)`

	return r.db.Exec(ctx, q, organization.Name, organization.Bin)
}

func (r *St) FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error) {
	q := `
		SELECT name, bin FROM organization`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	organizations := make([]entities.Organization, 0)

	for rows.Next() {
		var org entities.Organization

		err = rows.Scan(&org.Name, &org.Bin)
		if err != nil {
			return nil, r.db.HErr(err)
		}

		organizations = append(organizations, org)
	}
	if err = rows.Err(); err != nil {
		return nil, r.db.HErr(err)
	}

	return organizations, nil
}

func (r *St) FindOneOrganization(ctx context.Context, bin string) (entities.Organization, error) {
	q := `
		SELECT NAME, BIN FROM organization WHERE BIN = &1`

	var org entities.Organization

	err := r.db.QueryRow(ctx, q, bin).Scan(&org.Name, &org.Bin)
	if err != nil {
		err = r.db.HErr(err)
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

func (r *St) DeleteOrganization(ctx context.Context, bin string) error {
	q := `
		DELETE FROM organization
		WHERE bin = $1;`

	if err := r.db.Exec(ctx, q, bin); err != nil {
		return err
	}

	return nil
}
