package core

import (
	"kaspi-qr/internal/adapters/provider/kaspi"
)

func getKaspiStatusCodeDescription() map[int]string {
	return map[int]string{
		kaspi.SuccessStatus:                         "Успешный статус операции",
		kaspi.NoCertificateStatus:                   "Отсутствует сертификат клиента",
		kaspi.DeviceNotFoundStatus:                  "Устройство с заданным идентификатором не найдено",
		kaspi.DeviceDeactivatedStatus:               "Устройство не активно (отключено или удалено)",
		kaspi.DeviceAlreadyExistStatus:              "Устройство уже добавлено в другую торговую точку",
		kaspi.PurchaseNotFoundStatus:                "Покупка не найдена",
		kaspi.TradePointsDoesntExistStatus:          "Отсутствуют торговые точки, необходимо создать торговую точку в приложении Kaspi Pay",
		kaspi.TradePointNotFound:                    "Торговая точка не найдена",
		kaspi.RefundAmountGreaterStatus:             "Сумма возврата не может превышать сумму покупки",
		kaspi.RefundErrorStatus:                     "Ошибка возврата, необходимо попробовать еще раз и при повторении ошибки обратиться в банк",
		kaspi.TradePointDeactivatedStatus:           "Торговая точка отключена",
		kaspi.TradePointDoesntAcceptQrPaymentStatus: "Торговая точка не принимает оплату с QR",
		kaspi.WrongAmountStatus:                     "Указана неверная сумма операции",
		kaspi.NoPaymentMethodsAvailableStatus:       "Нет доступных методов оплаты",
		kaspi.PurchaseUuidNotFoundStatus:            "Покупка с заданным идентификатором не найдена",
		kaspi.TradePointDoesntMatchDeviceStatus:     "Торговая точка покупки не соответствует текущему устройству",
		kaspi.WrongPurchaseStatus:                   "Невозможно вернуть покупку (несоответствующий статус покупки)",
		kaspi.PartialRefundNotAvailableStatus:       "Частичный возврат невозможен",
		kaspi.ServiceNotAvailableStatus:             "Сервис временно недоступен",
	}
}

func (s *St) SetMessageByStatusCode(statusCode int) string {
	statusMessageMap := getKaspiStatusCodeDescription()

	return statusMessageMap[statusCode]
}
