package usecases

import (
	"context"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopTools"
)

func (u *St) OrganisationList(ctx context.Context,
	pars *entities.OrganisationListParsSt) ([]*entities.OrganisationSt, int64, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	if err = dopTools.RequirePageSize(pars.ListParams, cns.MaxPageSize); err != nil {
		return nil, 0, err
	}

	return u.cr.Organisation.List(ctx, pars)
}

func (u *St) OrganisationGet(ctx context.Context, id string) (*entities.OrganisationSt, error) {
	// var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return nil, 0, err
	// }

	return u.cr.Organisation.Get(ctx, id, true)
}

func (u *St) OrganisationCreate(ctx context.Context,
	obj *entities.OrganisationCUSt) (string, error) {
	var err error

	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return "", err
	// }

	var result string

	err = u.db.TransactionFn(ctx, func(ctx context.Context) error {
		result, err = u.cr.Organisation.Create(ctx, obj)
		return err
	})

	return result, err
}

func (u *St) OrganisationUpdate(ctx context.Context,
	id string, obj *entities.OrganisationCUSt) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Organisation.Update(ctx, id, obj)
	})
}

func (u *St) OrganisationDelete(ctx context.Context,
	id string) error {
	// ses := u.SessionGetFromContext(ctx)
	//
	// if err = u.SessionRequireAuth(ses); err != nil {
	// 	return err
	// }

	return u.db.TransactionFn(ctx, func(ctx context.Context) error {
		return u.cr.Organisation.Delete(ctx, id)
	})
}
