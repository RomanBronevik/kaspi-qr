package repo

import (
	"context"
	"kaspi-qr/internal/domain/entities"
)

type Repo interface {
	// organisation
	OrganisationGet(ctx context.Context, id string) (*entities.OrganisationSt, error)
	OrganisationList(ctx context.Context, pars *entities.OrganisationListParsSt) ([]*entities.OrganisationSt, int64, error)
	OrganisationIdExists(ctx context.Context, id string) (bool, error)
	OrganisationCreate(ctx context.Context, obj *entities.OrganisationCUSt) (string, error)
	OrganisationUpdate(ctx context.Context, id string, obj *entities.OrganisationCUSt) error
	OrganisationDelete(ctx context.Context, id string) error
}
