package entities

type CitySt struct {
	Id     string `json:"id" db:"id"`
	Code   string `json:"code" db:"code"`
	Name   string `json:"name" db:"name"`
	OrgBin string `json:"org_bin" db:"org_bin"`
}

type CityListParsSt struct {
	Ids    *[]string `json:"ids" form:"ids"`
	Code   *string   `json:"code" form:"code"`
	OrgBin *string   `json:"org_bin" form:"org_bin"`
}

type CityCUSt struct {
	Id     *string `json:"id" db:"id"`
	Code   *string `json:"code" db:"code"`
	Name   *string `json:"name" db:"name"`
	OrgBin *string `json:"org_bin" db:"org_bin"`
}
