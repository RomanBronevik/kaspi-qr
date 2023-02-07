package organization

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"kaspi-qr/internal/organization"
	"kaspi-qr/pkg/repository"
)

type repo struct {
	client repository.Client
}

func (r *repo) Create(ctx context.Context, organization *organization.Organization) error {
	q := `
		INSERT INTO organization (bin) 
		VALUES ($1) 
		RETURNING id`
	if err := r.client.QueryRow(ctx, q, organization.Bin).Scan(&organization.ID); err != nil {
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

func (r *repo) FindAll(ctx context.Context) (u []organization.Organization, err error) {
	q := `
		SELECT ID, BIN FROM public.organization`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	organizations := make([]organization.Organization, 0)

	for rows.Next() {
		var org organization.Organization

		err := rows.Scan(&org.ID, &org.Bin)
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

func (r *repo) FindOne(ctx context.Context, id string) (organization.Organization, error) {
	q := `
		SELECT ID, BIN FROM public.organization WHERE id = &1`

	//Trace

	var org organization.Organization
	err := r.client.QueryRow(ctx, q, id).Scan(&org.ID, &org.Bin)
	if err != nil {
		return organization.Organization{}, err
	}

	return org, nil

}

func (r *repo) Update(ctx context.Context, organization organization.Organization) error {
	//TODO implement me
	panic("implement me")
}

func (r *repo) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client repository.Client) *repo {
	return &repo{
		client: client,
	}
}
