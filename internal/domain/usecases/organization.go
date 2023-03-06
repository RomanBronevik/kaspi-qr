package usecases

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateOrganization(ctx *gin.Context, obj *entities.CreateOrganizationDTO) error {
	err := s.cr.CreateOrganization(ctx, obj)

	return err
}

//	func (s *St) UpdateOrganization(ctx *gin.Context, obj *entities.Organization) error {
//		err := s.cr.UpdateOrganization(ctx, obj)
//
//		return err
//	}
func (s *St) DeleteOrganization(ctx *gin.Context, bin string) error {
	err := s.cr.DeleteOrganization(ctx, bin)

	return err
}

func (s *St) FindAllOrganizations(ctx *gin.Context) ([]entities.Organization, error) {
	organizations, err := s.cr.FindAllOrganizations(ctx)

	return organizations, err
}

func (s *St) FindOneOrganization(ctx *gin.Context, obj *entities.Organization) (entities.Organization, error) {
	organization, err := s.cr.FindOneOrganization(ctx, obj)

	return organization, err
}
