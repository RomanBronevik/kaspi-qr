package entities

import "time"

type PaymentSt struct {
	Id              string       `json:"id" db:"id"`
	Created         time.Time    `json:"created" db:"created"`
	Modified        time.Time    `json:"modified" db:"modified"`
	OrdId           string       `json:"ord_id" db:"ord_id"`
	Status          string       `json:"status" db:"status"`
	StatusChangedAt time.Time    `json:"status_changed_at" db:"status_changed_at"`
	PaymentMethod   string       `json:"payment_method" db:"payment_method"`
	Amount          float64      `json:"amount" db:"amount"`
	ExpireDt        time.Time    `json:"expire_dt" db:"expire_dt"`
	Pbo             PaymentPboSt `json:"pbo" db:"pbo"`
}

type PaymentPboSt struct {
	StatusPollingInterval      int `json:"status_polling_interval"`
	LinkActivationWaitTimeout  int `json:"link_activation_wait_timeout"`
	PaymentConfirmationTimeout int `json:"payment_confirmation_timeout"`
}

type PaymentListParsSt struct {
	Ids           *[]string `json:"ids" form:"ids"`
	OrdId         *string   `json:"ord_id" form:"ord_id"`
	Status        *string   `json:"status" form:"status"`
	PaymentMethod *string   `json:"payment_method" form:"payment_method"`
}

type PaymentCUSt struct {
	Id              *string       `json:"id" db:"id"`
	Modified        *time.Time    `json:"-" db:"modified"`
	OrdId           *string       `json:"ord_id" db:"ord_id"`
	Status          *string       `json:"status" db:"status"`
	StatusChangedAt *time.Time    `json:"-" db:"status_changed_at"`
	PaymentMethod   *string       `json:"payment_method" db:"payment_method"`
	Amount          *float64      `json:"amount" db:"amount"`
	ExpireDt        *time.Time    `json:"expire_dt" db:"expire_dt"`
	Pbo             *PaymentPboSt `json:"pbo" db:"pbo"`
}
