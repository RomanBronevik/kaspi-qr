package usecases

import (
	"kaspi-qr/internal/domain/entities"

	"github.com/gin-gonic/gin"
)

func (u *St) CreateOrganization(ctx *gin.Context, obj *entities.CreateOrganizationDTO) error {
	err := u.cr.CreateOrganization(ctx, obj)

	return err
}

//	func (s *St) UpdateOrganization(ctx *gin.Context, obj *entities.Organization) error {
//		err := s.cr.UpdateOrganization(ctx, obj)
//
//		return err
//	}
func (u *St) DeleteOrganization(ctx *gin.Context, bin string) error {
	err := u.cr.DeleteOrganization(ctx, bin)

	return err
}

func (u *St) FindAllOrganizations(ctx *gin.Context) ([]entities.Organization, error) {
	organizations, err := u.cr.FindAllOrganizations(ctx)

	return organizations, err
}

func (u *St) FindOneOrganization(ctx *gin.Context, obj *entities.Organization) (entities.Organization, error) {
	organization, err := u.cr.FindOneOrganization(ctx, obj)

	return organization, err
}
