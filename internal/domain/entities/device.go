package entities

import "time"

type DeviceSt struct {
	Id           string    `json:"id" db:"id"`
	Created      time.Time `json:"created" db:"created"`
	Token        string    `json:"token" db:"token"`
	TradePointId int64     `json:"trade_point_id" db:"trade_point_id"`
	OrgBin       string    `json:"org_bin" db:"org_bin"`
}

type DeviceListParsSt struct {
	Ids          *[]string `json:"ids" form:"ids"`
	Token        *string   `json:"token" form:"token"`
	TradePointId *int64    `json:"trade_point_id" form:"trade_point_id"`
	OrgBin       *string   `json:"org_bin" form:"org_bin"`
	//CityId       *string   `json:"city_id" form:"city_id"`
}

type DeviceCUSt struct {
	Id           *string `json:"id" db:"id"`
	Token        *string `json:"-" db:"token"`
	TradePointId *int64  `json:"trade_point_id" db:"trade_point_id"`
	OrgBin       *string `json:"org_bin" db:"org_bin"`
}
