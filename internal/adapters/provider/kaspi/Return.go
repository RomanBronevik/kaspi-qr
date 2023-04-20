package kaspi

import (
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/errs"
	"strconv"
)

func (s *St) KaspiOperationDetails(reqObj provider.OperationGetReqSt) (*provider.OperationDetailsSt, error) {
	uriPath := "payment/details?QrPaymentId=" + strconv.FormatInt(reqObj.QrPaymentId, 10) + "&DeviceToken=" + reqObj.DeviceToken

	repObj := &provider.OperationGetRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("OperationDetails", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("OperationDetails bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) KaspiReturnWithoutClient(input provider.ReturnReqSt) (int64, error) {
	uriPath := "payment/return"

	repObj := &provider.ReturnRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, input, repObj)
	if err != nil {
		resp.LogError("ReturnWithoutClient", err)
		return 0, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("ReturnWithoutClient bad status-code", err)
		return 0, errs.ServiceNA
	}

	return repObj.ReturnOperationDataSt.ReturnOperationId, nil
}
