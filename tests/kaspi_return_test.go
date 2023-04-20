package tests

//
//import (
//	"encoding/json"
//	"github.com/spf13/viper"
//	"github.com/stretchr/testify/require"
//	"io"
//	"kaspi-qr/internal/adapters/provider/kaspi"
//	"kaspi-qr/internal/domain/entities"
//	"testing"
//)
//
//type Details struct {
//	QrPaymentId int
//	DeviceToken string
//}
//
//func TestOperationDetails(t *testing.T) {
//	ViperAndOsConfig()
//
//	input := Details{
//		QrPaymentId: 504710605,
//		DeviceToken: "0c2f8ec2-8c98-46d0-b3d9-06015e2fdba3",
//	}
//
//	reader, writer := io.Pipe()
//
//	go func() {
//		defer writer.Close()
//		err := json.NewEncoder(writer).Encode(input)
//		require.Nil(t, err)
//	}()
//
//	outputBody, err := kaspi.PaymentGetDetails(reader)
//	require.Nil(t, err)
//	require.Nil(t, outputBody.Data)
//	require.Equal(t, outputBody.StatusCode, kaspi.PurchaseNotFoundStatus)
//}
//
//func TestRefund(t *testing.T) {
//	ViperAndOsConfig()
//	input := entities.ReturnRequestInput{
//		DeviceToken:     "0c2f8ec2-8c98-46d0-b3d9-06015e2fdba3",
//		OrganizationBin: viper.GetString("testBin"),
//		QrPaymentId:     504709952,
//		Amount:          200,
//	}
//
//	outputBody, err := kaspi.PaymentReturn(input)
//	require.Nil(t, err)
//	require.Equal(t, outputBody.StatusCode, kaspi.WrongPurchaseStatus)
//}
