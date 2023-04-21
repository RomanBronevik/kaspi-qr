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
	OrsStatusCreated       = "created"
	OrsStatusLinkActivated = "link_activated"
	OrsStatusPaid          = "paid"
	OrsStatusError         = "error"
	OrsStatusRefunded      = "refunded"
)

func OrsStatusIsValid(v string) bool {
	return v == OrsStatusCreated ||
		v == OrsStatusLinkActivated ||
		v == OrsStatusPaid ||
		v == OrsStatusError ||
		v == OrsStatusRefunded
}

const (
	PaymentStatusCreated       = "created"
	PaymentStatusLinkActivated = "link_activated"
	PaymentStatusPaid          = "paid"
	PaymentStatusError         = "error"
	PaymentStatusRefunded      = "refunded"
)

func PaymentStatusIsValid(v string) bool {
	return v == PaymentStatusCreated ||
		v == PaymentStatusLinkActivated ||
		v == PaymentStatusPaid ||
		v == PaymentStatusError ||
		v == PaymentStatusRefunded
}
