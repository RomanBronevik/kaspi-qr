package kaspi

import (
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/errs"
)

func (s *St) CreateQrToken(reqObj provider.PaymentCreateReqSt) (*provider.PaymentSt, error) {
	uriPath := "qr/create"

	repObj := &provider.PaymentCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("CreateQrToken", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("CreateQrToken bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) CreatePaymentLink(reqObj provider.PaymentLinkCreateReqSt) (*provider.PaymentLinkSt, error) {
	uriPath := "qr/create-link"

	repObj := &provider.PaymentLinkCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("CreatePaymentLink", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("CreatePaymentLink bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) OperationStatus(qrPaymentId string) (string, error) {
	uriPath := "payment/status/" + qrPaymentId

	repObj := &provider.OperationStatusRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("OperationStatus", err)
		return "", err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("OperationStatus bad status-code", err)
		return "", errs.ServiceNA
	}

	return repObj.Data.Status, nil
}
