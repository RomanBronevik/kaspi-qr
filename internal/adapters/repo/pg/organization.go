package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error {
	q := `
		INSERT INTO organization (name, bin) 
		VALUES ($1, $2)`

	if _, err := r.client.Exec(ctx, q, organization.Name, organization.Bin); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error) {
	q := `
		SELECT NAME, BIN FROM public.organization`
	rows, err := r.client.Query(ctx, q)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	organizations := make([]entities.Organization, 0)

	for rows.Next() {
		var org entities.Organization

		err := rows.Scan(&org.Name, &org.Bin)
		if err != nil {
			return nil, err
		}

		organizations = append(organizations, org)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return organizations, nil
}

func (r *St) FindOneOrganization(ctx context.Context, bin string) (entities.Organization, error) {
	q := `
		SELECT NAME, BIN FROM public.organization WHERE BIN = &1`

	//Trace

	var org entities.Organization
	err := r.client.QueryRow(ctx, q, bin).Scan(&org.Name, &org.Bin)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.Organization{}, err
	}

	return org, nil

}

//func (r *St) UpdateOrganization(ctx context.Context, organization entities.Organization) error {
//	//TODO implement me
//	panic("implement me")
//}

func (r *St) DeleteOrganization(ctx context.Context, bin string) error {
	q := `
		DELETE FROM organization
		WHERE bin = $1;`

	if _, err := r.client.Exec(ctx, q, bin); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return r.ErorrHandler(err)
	}

	return nil
}
