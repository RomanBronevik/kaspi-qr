package kaspi

import (
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/errs"
)

func (s *St) GetAllTradePoints(orgBin string) ([]*provider.TradePointSt, error) {
	uriPath := "partner/tradepoints/" + orgBin

	repObj := &provider.TradePointListRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("GetAllTradePoints", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("DeviceDelete bad status-code", err)
		return nil, errs.ServiceNA
	}

	return repObj.Data, nil
}

func (s *St) DeviceRegistration(reqObj provider.DeviceCreateReqSt) (string, error) {
	uriPath := "device/register/"

	repObj := &provider.DeviceCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("DeviceRegistration", err)
		return "", err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("DeviceDelete bad status-code", err)
		return "", errs.ServiceNA
	}

	return repObj.Data.DeviceToken, nil
}

func (s *St) DeviceDelete(reqObj provider.DeviceRemoveReqSt) error {
	uriPath := "device/delete/"

	repObj := &provider.BaseRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("DeviceDelete", err)
		return err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("DeviceDelete bad status-code", err)
		return errs.ServiceNA
	}

	return nil
}
