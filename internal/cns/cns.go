package cns

const UnInteger = 1000

const (
	StatusCreated   = "Created"
	StatusWait      = "Wait"
	StatusProcessed = "Processed"
	StatusSuccess   = "Success"
	StatusError     = "Error"
	StatusRefund    = "Refunded"
)

const (
	PaymentMethodQr   = "Qr"
	PaymentMethodLink = "PaymentLink"
)

const HoursQuantity = 72

const (
	OrdSrcSite = "site"
)

func OrdSrcIsValid(v string) bool {
	return v == OrdSrcSite
}

const (
	PlatformSite    = "site"
	PlatformMSite   = "m_site"
	PlatformIOS     = "ios"
	PlatformAndroid = "android"
)

func PlatformIsValid(v string) bool {
	return v == PlatformSite ||
		v == PlatformMSite ||
		v == PlatformIOS ||
		v == PlatformAndroid
}
