package entities

import "time"

type PaymentSt struct {
	Id              int64        `json:"id" db:"id"`
	Created         time.Time    `json:"created" db:"created"`
	Modified        time.Time    `json:"modified" db:"modified"`
	OrdId           string       `json:"ord_id" db:"ord_id"`
	Link            string       `json:"link" db:"link"`
	Status          string       `json:"status" db:"status"`
	StatusChangedAt time.Time    `json:"status_changed_at" db:"status_changed_at"`
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
	Ids      *[]int64  `json:"ids" form:"ids"`
	OrdId    *string   `json:"ord_id" form:"ord_id"`
	Status   *string   `json:"status" form:"status"`
	Statuses *[]string `json:"statuses" form:"statuses"`
}

type PaymentCUSt struct {
	Id              *int64        `json:"id" db:"id"`
	Modified        *time.Time    `json:"-" db:"modified"`
	OrdId           *string       `json:"ord_id" db:"ord_id"`
	Link            *string       `json:"link" db:"link"`
	Status          *string       `json:"status" db:"status"`
	StatusChangedAt *time.Time    `json:"-" db:"status_changed_at"`
	Amount          *float64      `json:"amount" db:"amount"`
	ExpireDt        *time.Time    `json:"expire_dt" db:"expire_dt"`
	Pbo             *PaymentPboSt `json:"pbo" db:"pbo"`
}
