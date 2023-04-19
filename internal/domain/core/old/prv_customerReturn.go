package old

import "kaspi-qr/internal/domain/entities"

func (s *St) OperationDetails(input entities.OperationGetSt) (entities.OperationDetails, error) {
	output, err := s.kaspi.KaspiOperationDetails(input)

	return output, err
}

func (s *St) ReturnWithoutClient(input entities.ReturnRequestInput) (entities.ReturnSt, error) {
	output, err := s.kaspi.KaspiReturnWithoutClient(input)

	return output, err
}
