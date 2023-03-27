package usecases

import (
	"kaspi-qr/internal/domain/entities"
)

func (u *St) OperationDetails(input entities.OperationGetSt) (entities.OperationDetails, error) {
	output, err := u.cr.OperationDetails(input)

	return output, err
}

func (u *St) ReturnWithoutClient(input entities.ReturnRequestInput) (entities.ReturnSt, error) {
	output, err := u.cr.ReturnWithoutClient(input)

	return output, err
}
