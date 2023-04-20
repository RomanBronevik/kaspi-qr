package tests

//
//import (
//	"github.com/spf13/viper"
//	"github.com/stretchr/testify/require"
//	"kaspi-qr/internal/adapters/provider/kaspi"
//	"kaspi-qr/internal/domain/entities"
//	"testing"
//)
//
//func TestDeviceRegistration(t *testing.T) {
//	ViperAndOsConfig()
//
//	requestBody := entities.DeviceInputReg{
//		DeviceId:        "GFC-4563928",
//		OrganizationBin: viper.GetString("testBin"),
//		TradePointId:    "19",
//	}
//
//	outputBody, err := kaspi.DeviceCreate(requestBody)
//	require.Nil(t, err)
//	require.NotEmpty(t, outputBody.Data)
//	require.Equal(t, 0, outputBody.StatusCode)
//
//	deviceToken := entities.DeviceToken{
//		DeviceToken: "23db7f5d-a863-46e3-8a12-41abf917013d",
//	}
//
//	equal := entities.DeviceOutputReg{
//		Data:       deviceToken,
//		Message:    "",
//		StatusCode: kaspi.SuccessStatus,
//	}
//
//	require.Equal(t, &equal, &outputBody)
//}
//
//func TestDeviceDelete(t *testing.T) {
//	ViperAndOsConfig()
//
//	requestBody := entities.DeviceInputDel{
//		OrganizationBin: viper.GetString("testBin"),
//		DeviceToken:     "9c1da238-eb56-4761-9fbf-8e3d0efd91ec",
//	}
//
//	outputBody, err := kaspi.DeviceDelete(requestBody)
//	require.Nil(t, err)
//	require.Equal(t, kaspi.DeviceNotFoundStatus, outputBody.StatusCode)
//
//	equal := entities.DeviceOutputDel{
//		StatusCode: kaspi.DeviceNotFoundStatus,
//		Message:    "",
//	}
//
//	require.Equal(t, &equal, &outputBody)
//
//	requestBody.DeviceToken = "5f288da2-611b-4e04-90c5-85a48a16f03b"
//
//	outputBody, err = kaspi.DeviceDelete(requestBody)
//	require.Nil(t, err)
//	require.Equal(t, kaspi.SuccessStatus, outputBody.StatusCode)
//
//	equal = entities.DeviceOutputDel{
//		StatusCode: kaspi.SuccessStatus,
//		Message:    "",
//	}
//
//	require.Equal(t, &equal, &outputBody)
//}
//
//func TestGetTradePoints(t *testing.T) {
//	ViperAndOsConfig()
//
//	tradePoints, err := kaspi.TradePointList(viper.GetString("testBin"))
//	require.Nil(t, err)
//	require.NotEmpty(t, tradePoints.Data)
//	require.Equal(t, 10, len(tradePoints.Data))
//	//var arrTradePoints = []entities.KaspiTradePointSt{}
//	//
//	//arrTradePoints[0] = entities.KaspiTradePointSt{
//	//	TradePointId: 104039,
//	//	TradePointName: "@beauty_zone_liliya (г. Алматы, улица Наурызбай Батыра, 154А)",
//	//}
//	//
//	//arrTradePoints[1] = entities.KaspiTradePointSt{
//	//	TradePointId: 104248,
//	//	TradePointName: "12345 (г. Алматы, проспект Абая, 59)",
//	//}
//	//
//	//arrTradePoints[2] = entities.KaspiTradePointSt{
//	//	TradePointId: 104251,
//	//	TradePointName: "aaaaaa (г. Алматы, проспект Абая, 28)",
//	//}
//	//
//	//arrTradePoints[3] = entities.KaspiTradePointSt{
//	//	TradePointId: 104246,
//	//	TradePointName: "Hhhh (г. Алматы, проспект Абая, 28с2)",
//	//}
//	//
//	//arrTradePoints[4] = entities.KaspiTradePointSt{
//	//	TradePointId: 19,
//	//	TradePointName: "IPlaza (электроника) (г. Алматы, ул. Абылай хана, 3А)",
//	//}
//	//
//	//arrTradePoints[5] = entities.KaspiTradePointSt{
//	//	TradePointId: 104263,
//	//	TradePointName: "Kaspi (г. Алматы, ул. Абай, 6)",
//	//}
//	//
//	//arrTradePoints[6] = entities.KaspiTradePointSt{
//	//	TradePointId: 104533,
//	//	TradePointName: "Kaspi (г. Алматы, ул. Абылай хан, 48)",
//	//}
//	//
//	//arrTradePoints[7] = entities.KaspiTradePointSt{
//	//	TradePointId: 104250,
//	//	TradePointName: "kkkkkkkk (г. Алматы, улица Желтоксан, 166)",
//	//}
//	//
//	//arrTradePoints[8] = entities.KaspiTradePointSt{
//	//	TradePointId: 104247,
//	//	TradePointName: "MDM (г. Алматы, улица Наурызбай Батыра, 154А)",
//	//}
//	//
//	//arrTradePoints[9] = entities.KaspiTradePointSt{
//	//	TradePointId: 104176,
//	//	TradePointName: "RoxPoint (г. Алматы, улица Бекхожина, 5А)",
//	//}
//
//	//equal := entities.TradePointSt{
//	//	StatusCode: kaspi.SuccessStatus,
//	//	Message: "Успешный статус операции",
//	//	Data: arrTradePoints,
//	//}
//
//}
