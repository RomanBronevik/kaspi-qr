package core

import (
	"github.com/gin-gonic/gin"
	"kaspi-qr/internal/domain/entities"
)

func (s *St) CreateOrganization(ctx *gin.Context, obj *entities.CreateOrganizationDTO) error {
	err := s.repo.CreateOrganization(ctx, obj)

	return err
}

func (s *St) DeleteOrganization(ctx *gin.Context, bin string) error {
	err := s.repo.DeleteOrganization(ctx, bin)

	return err
}

//
//func (s *St) UpdateOrganization(ctx *gin.Context, obj *entities.Organization) error {
//	err := s.repo.UpdateOrganization(ctx, obj)
//
//	return err
//}

func (s *St) FindAllOrganizations(ctx *gin.Context) ([]entities.Organization, error) {
	organizations, err := s.repo.FindAllOrganizations(ctx)

	return organizations, err
}

func (s *St) FindOneOrganization(ctx *gin.Context, obj *entities.Organization) (entities.Organization, error) {
	organization, err := s.repo.FindOneOrganization(ctx, obj.Bin)

	return organization, err
}
