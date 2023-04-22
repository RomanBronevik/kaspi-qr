package entities

import "time"

type OrdSt struct {
	Id       string    `json:"id" db:"id"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
	Src      string    `json:"src" db:"src"`
	DeviceId string    `json:"device_id" db:"device_id"`
	CityId   string    `json:"city_id" db:"city_id"`
	Amount   float64   `json:"amount" db:"amount"`
	Status   string    `json:"status" db:"status"`
	Platform string    `json:"platform" db:"platform"`
}

type OrdListParsSt struct {
	Ids      *[]string `json:"ids" form:"ids"`
	Src      *string   `json:"src" form:"src"`
	DeviceId *string   `json:"device_id" form:"device_id"`
	CityId   *string   `json:"city_id" form:"city_id"`
	Status   *string   `json:"status" form:"status"`
	Platform *string   `json:"platform" form:"platform"`
}

type OrdCUSt struct {
	Id       *string    `json:"id" db:"id"`
	Modified *time.Time `json:"-" db:"modified"`
	Src      *string    `json:"src" db:"src"`
	DeviceId *string    `json:"-" db:"device_id"`
	CityCode *string    `json:"city_code" db:"-"`
	CityId   *string    `json:"-" db:"city_id"`
	Amount   *float64   `json:"amount" db:"amount"`
	Status   *string    `json:"-" db:"status"`
	Platform *string    `json:"platform" db:"platform"`
}

type OrdCreateRepSt struct {
	PaymentId int64  `json:"payment_id"`
	QrUrl     string `json:"qr_url"`
	QrCode    string `json:"qr_code"`
}
