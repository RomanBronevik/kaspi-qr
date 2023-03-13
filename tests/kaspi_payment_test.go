package tests

//
//import (
//	"github.com/spf13/viper"
//	"github.com/stretchr/testify/require"
//	"kaspi-qr/internal/adapters/provider/kaspi"
//	"kaspi-qr/internal/domain/entities"
//	"strconv"
//	"testing"
//)
//
//func TestQrTokenGeneration(t *testing.T) {
//	ViperAndOsConfig()
//
//	qrToken := entities.KaspiPaymentInput{
//		OrganizationBin: viper.GetString("testBin"),
//		DeviceToken:     "0c2f8ec2-8c98-46d0-b3d9-06015e2fdba3",
//		Amount:          200,
//		ExternalId:      "123",
//	}
//
//	outputBody, err := kaspi.CreateQrToken(qrToken)
//	require.Nil(t, err)
//	require.Equal(t, kaspi.SuccessStatus, outputBody.StatusCode)
//	require.NotEmpty(t, outputBody.Data)
//	require.Len(t, strconv.Itoa(outputBody.Data.QrPaymentId), 9)
//}
//
//func TestPaymentLinkGeneration(t *testing.T) {
//	ViperAndOsConfig()
//
//	qrToken := entities.KaspiPaymentInput{
//		OrganizationBin: viper.GetString("testBin"),
//		DeviceToken:     "0c2f8ec2-8c98-46d0-b3d9-06015e2fdba3",
//		Amount:          200,
//		ExternalId:      "123",
//	}
//
//	outputBody, err := kaspi.CreatePaymentLink(qrToken)
//	require.Nil(t, err)
//	require.Equal(t, kaspi.SuccessStatus, outputBody.StatusCode)
//	require.NotEmpty(t, outputBody.Data)
//	require.Len(t, strconv.Itoa(outputBody.Data.PaymentId), 9)
//}
//
//func TestOperationStatus(t *testing.T) {
//	ViperAndOsConfig()
//	QrPaymentId := strconv.Itoa(504710605)
//	outputBody, err := kaspi.OperationStatus(QrPaymentId)
//	require.Nil(t, err)
//	require.Equal(t, kaspi.PurchaseNotFoundStatus, outputBody.StatusCode)
//	require.Nil(t, outputBody.Data)
//}
