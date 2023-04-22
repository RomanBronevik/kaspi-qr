package errs

// Err

type Err string

func (e Err) Error() string {
	return string(e)
}

// ErrWithDesc

type ErrWithDesc struct {
	Err  Err
	Desc string
}

func (e ErrWithDesc) Error() string {
	return e.Err.Error() + ", desc:" + e.Desc
}

// errors

const (
	BadJson          = Err("bad_json")
	BadQueryParams   = Err("bad_query_params")
	ServiceNA        = Err("service_not_available")
	NotImplemented   = Err("not_implemented")
	NotAuthorized    = Err("not_authorized")
	PermissionDenied = Err("permission_denied")
	ObjectNotFound   = Err("object_not_found")
	BadStatusCode    = Err("bad_status_code")

	IdRequired           = Err("id_required")
	IdTooLong            = Err("id_too_long")
	OrgBinRequired       = Err("org_bin_required")
	OrgBinTooLong        = Err("org_bin_too_long")
	TradePointIdRequired = Err("trade_point_id_required")
	SrcRequired          = Err("src_required")
	BadSrc               = Err("bad_src")
	CityCodeRequired     = Err("city_code_required")
	CityNotFound         = Err("city_not_found")
	AmountRequired       = Err("amount_required")
	AmountMustBePositive = Err("amount_must_be_positive")
	PlatformRequired     = Err("platform_required")
	BadPlatform          = Err("bad_platform")
	DeviceNotFound       = Err("device_not_found")
	OrderAlreadyPaid     = Err("order_already_paid")
	OrderIdRequired      = Err("order_id_required")
	OrderNotFound        = Err("order_not_found")
)
