package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateOrganization(ctx context.Context, organization *entities.CreateOrganizationDTO) error {
	q := `
		INSERT INTO organization (name, bin) 
		VALUES ($1, $2) 
		RETURNING id`

	var id string

	if err := r.client.QueryRow(ctx, q, organization.Name, organization.Bin).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			fmt.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *St) FindAllOrganizations(ctx context.Context) (u []entities.Organization, err error) {
	q := `
		SELECT ID, NAME, BIN FROM public.organization`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	organizations := make([]entities.Organization, 0)

	for rows.Next() {
		var org entities.Organization

		err := rows.Scan(&org.ID, &org.Name, &org.Bin)
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
		SELECT ID, NAME, BIN FROM public.organization WHERE BIN = &1`

	//Trace

	var org entities.Organization
	err := r.client.QueryRow(ctx, q, bin).Scan(&org.ID, &org.Name, &org.Bin)
	if err != nil {
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

	var id string

	if err := r.client.QueryRow(ctx, q, bin).Scan(&id); err != nil {
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
