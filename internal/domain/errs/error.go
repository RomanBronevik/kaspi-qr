package errs

import "github.com/rendau/dop/dopErrs"

// errors

const (
	IdRequired           = dopErrs.Err("id_required")
	IdTooLong            = dopErrs.Err("id_too_long")
	OrgBinRequired       = dopErrs.Err("org_bin_required")
	OrgBinTooLong        = dopErrs.Err("org_bin_too_long")
	TradePointIdRequired = dopErrs.Err("trade_point_id_required")
	SrcRequired          = dopErrs.Err("src_required")
	BadSrc               = dopErrs.Err("bad_src")
	CityCodeRequired     = dopErrs.Err("city_code_required")
	CityNotFound         = dopErrs.Err("city_not_found")
	AmountRequired       = dopErrs.Err("amount_required")
	AmountMustBePositive = dopErrs.Err("amount_must_be_positive")
	PlatformRequired     = dopErrs.Err("platform_required")
	BadPlatform          = dopErrs.Err("bad_platform")
	DeviceNotFound       = dopErrs.Err("device_not_found")
	OrderAlreadyPaid     = dopErrs.Err("order_already_paid")
	OrderIdRequired      = dopErrs.Err("order_id_required")
	OrderNotFound        = dopErrs.Err("order_not_found")
)
