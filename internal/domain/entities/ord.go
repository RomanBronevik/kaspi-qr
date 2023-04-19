package entities

import "time"

type OrdSt struct {
	Id       string    `json:"id" db:"id"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
	OrgBin   string    `json:"org_bin" db:"org_bin"`
	Status   string    `json:"status" db:"status"`
}

type OrdListParsSt struct {
	Ids    *[]string `json:"ids" form:"ids"`
	OrgBin *string   `json:"org_bin" form:"org_bin"`
	Status *string   `json:"status" form:"status"`
}

type OrdCUSt struct {
	Id       *string    `json:"id" db:"id"`
	Modified *time.Time `json:"-" db:"modified"`
	OrgBin   *string    `json:"org_bin" db:"org_bin"`
	Status   *string    `json:"status" db:"status"`
}
