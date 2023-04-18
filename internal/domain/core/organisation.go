package core

import (
	"context"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopErrs"
)

type Organisation struct {
	r *St
}

func NewOrganisation(r *St) *Organisation {
	return &Organisation{r: r}
}

func (c *Organisation) ValidateCU(ctx context.Context, obj *entities.OrganisationCUSt, id string) error {
	// forCreate := id == ""

	return nil
}

func (c *Organisation) List(ctx context.Context, pars *entities.OrganisationListParsSt) ([]*entities.OrganisationSt, int64, error) {
	items, tCount, err := c.r.repo.OrganisationList(ctx, pars)
	if err != nil {
		return nil, 0, err
	}

	return items, tCount, nil
}

func (c *Organisation) Get(ctx context.Context, id string, errNE bool) (*entities.OrganisationSt, error) {
	result, err := c.r.repo.OrganisationGet(ctx, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		if errNE {
			return nil, dopErrs.ObjectNotFound
		}
		return nil, nil
	}

	return result, nil
}

func (c *Organisation) IdExists(ctx context.Context, id string) (bool, error) {
	return c.r.repo.OrganisationIdExists(ctx, id)
}

func (c *Organisation) Create(ctx context.Context, obj *entities.OrganisationCUSt) (string, error) {
	var err error

	err = c.ValidateCU(ctx, obj, "")
	if err != nil {
		return "", err
	}

	// create
	result, err := c.r.repo.OrganisationCreate(ctx, obj)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Organisation) Update(ctx context.Context, id string, obj *entities.OrganisationCUSt) error {
	var err error

	err = c.ValidateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	err = c.r.repo.OrganisationUpdate(ctx, id, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Organisation) Delete(ctx context.Context, id string) error {
	return c.r.repo.OrganisationDelete(ctx, id)
}
