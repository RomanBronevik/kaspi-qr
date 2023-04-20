package kaspi

import (
	"kaspi-qr/internal/adapters/logger"
	"kaspi-qr/internal/adapters/provider"
	"kaspi-qr/internal/domain/errs"
	"strconv"
)

type St struct {
	lg         logger.Full
	httpClient *httpClientSt
}

func New(lg logger.Full, kaspiUrl, certPath, certPassword string) (*St, error) {
	httpClient, err := newHttpClient(lg, kaspiUrl, certPath, certPassword)
	if err != nil {
		lg.Errorw("newHttpClient", err)
		return nil, err
	}

	return &St{
		lg:         lg,
		httpClient: httpClient,
	}, nil
}

// TRADE POINT

func (s *St) TradePointList(orgBin string) ([]*provider.TradePointSt, error) {
	uriPath := "partner/tradepoints/" + orgBin

	repObj := &provider.TradePointListRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("TradePointList", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("TradePointList bad status-code", err)
		return nil, errs.ServiceNA
	}

	return repObj.Data, nil
}

// DEVICE

func (s *St) DeviceCreate(reqObj provider.DeviceCreateReqSt) (string, error) {
	uriPath := "device/register/"

	repObj := &provider.DeviceCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("DeviceCreate", err)
		return "", err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("DeviceCreate bad status-code", err)
		return "", errs.ServiceNA
	}

	return repObj.Data.DeviceToken, nil
}

func (s *St) DeviceDelete(reqObj provider.DeviceDeleteReqSt) error {
	uriPath := "device/delete"

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

// PAYMENT

func (s *St) PaymentCreate(reqObj provider.PaymentCreateReqSt) (*provider.PaymentSt, error) {
	uriPath := "qr/create"

	repObj := &provider.PaymentCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("PaymentCreate", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("PaymentCreate bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) PaymentLinkCreate(reqObj provider.PaymentLinkCreateReqSt) (*provider.PaymentLinkSt, error) {
	uriPath := "qr/create-link"

	repObj := &provider.PaymentLinkCreateRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("PaymentLinkCreate", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("PaymentLinkCreate bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) PaymentGetStatus(paymentId string) (string, error) {
	uriPath := "payment/status/" + paymentId

	repObj := &provider.PaymentStatusRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("PaymentGetStatus", err)
		return "", err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("PaymentGetStatus bad status-code", err)
		return "", errs.ServiceNA
	}

	return repObj.Data.Status, nil
}

func (s *St) PaymentGetDetails(paymentId int64, deviceToken string) (*provider.PaymentDetailsSt, error) {
	uriPath := "payment/details?QrPaymentId=" + strconv.FormatInt(paymentId, 10) + "&DeviceToken=" + deviceToken

	repObj := &provider.PaymentDetailsRepSt{}

	resp, err := s.httpClient.sendRequest("GET", uriPath, nil, repObj)
	if err != nil {
		resp.LogError("PaymentDetails", err)
		return nil, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("PaymentDetails bad status-code", err)
		return nil, errs.ServiceNA
	}

	return &repObj.Data, nil
}

func (s *St) PaymentReturn(reqObj provider.PaymentReturnReqSt) (int64, error) {
	uriPath := "payment/return"

	repObj := &provider.PaymentReturnRepSt{}

	resp, err := s.httpClient.sendRequest("POST", uriPath, reqObj, repObj)
	if err != nil {
		resp.LogError("PaymentReturn", err)
		return 0, err
	}

	if repObj.StatusCode != SuccessStatus {
		resp.LogError("PaymentReturn bad status-code", err)
		return 0, errs.ServiceNA
	}

	return repObj.ReturnOperationDataSt.ReturnOperationId, nil
}
