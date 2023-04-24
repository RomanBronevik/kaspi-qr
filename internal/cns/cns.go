package cns

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

const (
	OrdStatusCreated  = "created"
	OrdStatusPaid     = "paid"
	OrdStatusError    = "error"
	OrdStatusExpired  = "expired"
	OrdStatusRefunded = "refunded"
)

const (
	PaymentStatusCreated       = "created"
	PaymentStatusLinkActivated = "link_activated"
	PaymentStatusPaid          = "paid"
	PaymentStatusError         = "error"
	PaymentStatusExpired       = "expired"
	PaymentStatusRefunded      = "refunded"
)
