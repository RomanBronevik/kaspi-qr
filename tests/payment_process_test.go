package tests

import (
	"crypto/rand"
	"errors"
	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/domain/entities"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	ViperAndOsConfig()

	requestBody := entities.DeviceInputReg{
		DeviceId:        "GFC-4563928",
		OrganizationBin: viper.GetString("testBin"),
		TradePointId:    "19",
	}

	outputBody, err := DeviceRegistration(requestBody)
	require.Nil(t, err)
	require.NotEmpty(t, outputBody.Data)
	require.Equal(t, 0, outputBody.StatusCode)

	qrToken := entities.KaspiPaymentInput{
		OrganizationBin: viper.GetString("testBin"),
		DeviceToken:     outputBody.Data.DeviceToken,
		Amount:          200,
		ExternalId:      "123",
	}

	outputBodyQr, err := CreateQrToken(qrToken)
	require.Nil(t, err)
	require.Equal(t, kaspi.StatusSuccess, outputBody.StatusCode)
	require.NotEmpty(t, outputBody.Data)
	require.Len(t, strconv.Itoa(outputBodyQr.Data.QrPaymentId), 9)

	PaymentId := strconv.Itoa(outputBodyQr.Data.QrPaymentId)

	outputBodyOp, err := OperationStatus(PaymentId)
	require.Nil(t, err)
	require.Equal(t, kaspi.StatusSuccess, outputBodyOp.StatusCode)

	refundSt := entities.ReturnRequestInput{
		QrPaymentId:     outputBodyQr.Data.QrPaymentId,
		OrganizationBin: viper.GetString("testBin"),
		DeviceToken:     outputBody.Data.DeviceToken,
		Amount:          200,
	}
	output, err := app.kaspi.PaymentReturn(refundSt)
	require.Nil(t, err)
	require.Equal(t, kaspi.PurchaseNotFoundStatus, output.StatusCode)

}

func TestProcessLink(t *testing.T) {
	ViperAndOsConfig()

	requestBody := entities.DeviceInputReg{
		DeviceId:        "GFC-4563928",
		OrganizationBin: viper.GetString("testBin"),
		TradePointId:    "19",
	}

	outputBody, err := DeviceRegistration(requestBody)
	require.Nil(t, err)
	require.NotEmpty(t, outputBody.Data)
	require.Equal(t, 0, outputBody.StatusCode)

	paymentLink := entities.KaspiPaymentInput{
		OrganizationBin: viper.GetString("testBin"),
		DeviceToken:     outputBody.Data.DeviceToken,
		Amount:          200,
		ExternalId:      "123",
	}

	outputBodyPayment, err := CreatePaymentLink(paymentLink)
	require.Nil(t, err)
	require.Equal(t, kaspi.StatusSuccess, outputBody.StatusCode)
	require.NotEmpty(t, outputBody.Data)
	require.Len(t, strconv.Itoa(outputBodyPayment.Data.PaymentId), 9)

	PaymentId := strconv.Itoa(outputBodyPayment.Data.PaymentId)

	outputBodyOp, err := OperationStatus(PaymentId)
	require.Nil(t, err)
	require.Equal(t, kaspi.StatusSuccess, outputBodyOp.StatusCode)

	//

	refundSt := entities.ReturnRequestInput{
		QrPaymentId:     outputBodyPayment.Data.PaymentId,
		OrganizationBin: viper.GetString("testBin"),
		DeviceToken:     outputBody.Data.DeviceToken,
		Amount:          200,
	}
	output, err := app.kaspi.PaymentReturn(refundSt)
	require.Nil(t, err)
	require.Equal(t, kaspi.PurchaseNotFoundStatus, output.StatusCode)

}

func DeviceRegistration(input entities.DeviceInputReg) (entities.DeviceOutputReg, error) {
	if len(input.DeviceId) == 0 || len(input.TradePointId) == 0 || len(input.OrganizationBin) == 0 {
		output := entities.DeviceOutputReg{
			StatusCode: kaspi.DeviceDeactivatedStatus,
			Message:    "Good",
			Data:       nil,
		}
		return output, errors.New("empty fields")
	}

	output := entities.DeviceOutputReg{
		StatusCode: 0,
		Message:    "Good",
		Data: &entities.DeviceToken{
			DeviceToken: "cd10060c-1888-4617-9216-8be4b7e46963",
		},
	}

	return output, nil
}

func CreateQrToken(input entities.KaspiPaymentInput) (entities.QrTokenOutput, error) {
	if len(input.ExternalId) == 0 || len(input.DeviceToken) == 0 || len(input.OrganizationBin) == 0 || input.Amount == 0 {
		output := entities.QrTokenOutput{
			StatusCode: kaspi.DeviceDeactivatedStatus,
			Message:    "Bad",
			Data:       nil,
		}
		return output, errors.New("empty fields")
	}

	output := entities.QrTokenOutput{
		StatusCode: kaspi.StatusSuccess,
		Message:    "Good",
		Data: &entities.QRStruct{
			QRToken:        "57293259759247999078793654586228299243435",
			ExpireDate:     time.Now().Local(),
			QrPaymentId:    RandomPaymentId(),
			PaymentMethods: []string{"Loan", "Gold", "Red"},
			QrPaymentBehaviorOptions: &entities.QrPaymentBehaviorOptions{
				StatusPollingInterval:      5,
				QrCodeScanWaitTimeout:      180,
				PaymentConfirmationTimeout: 65,
			},
		},
	}

	return output, nil
}

func RandomPaymentId() int {
	num, _ := rand.Int(rand.Reader, big.NewInt(1e9-1e8))
	num = num.Add(num, big.NewInt(1e8))
	return int(num.Int64())
}

func CreatePaymentLink(input entities.KaspiPaymentInput) (entities.PaymentLinkRequestOutput, error) {
	if len(input.ExternalId) == 0 || len(input.DeviceToken) == 0 || len(input.OrganizationBin) == 0 || input.Amount == 0 {
		output := entities.PaymentLinkRequestOutput{
			StatusCode: kaspi.DeviceDeactivatedStatus,
			Message:    "Bad",
			Data:       nil,
		}
		return output, errors.New("empty fields")
	}

	output := entities.PaymentLinkRequestOutput{
		StatusCode: kaspi.StatusSuccess,
		Message:    "Good",
		Data: &entities.PaymentLinkSt{
			PaymentLink:    "57293259759247999078793654586228299243435",
			ExpireDate:     time.Now().Local(),
			PaymentId:      RandomPaymentId(),
			PaymentMethods: []string{"Loan", "Gold", "Red"},
			PaymentBehaviorOptions: &entities.PaymentBehaviorOptions{
				StatusPollingInterval:      5,
				LinkActivationWaitTimeout:  180,
				PaymentConfirmationTimeout: 65,
			},
		},
	}

	return output, nil
}

func OperationStatus(PaymentId string) (entities.OperationStatus, error) {
	if len(PaymentId) != 9 {
		output := entities.OperationStatus{
			StatusCode: kaspi.PurchaseNotFoundStatus,
			Message:    "Bad",
			Data:       nil,
		}

		return output, errors.New("empty fields")
	}

	output := entities.OperationStatus{
		StatusCode: kaspi.StatusSuccess,
		Message:    "Good",
		Data: &entities.StatusSt{
			Status: kaspi.PaymentStatusProcessed,
		},
	}

	return output, nil
}
