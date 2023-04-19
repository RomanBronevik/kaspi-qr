package entities

type DeviceSt struct {
	Id     string `json:"id" db:"id"`
	Token  string `json:"token" db:"token"`
	OrgBin string `json:"org_bin" db:"org_bin"`
}

type DeviceListParsSt struct {
	Ids    *[]string `json:"ids" form:"ids"`
	Token  *string   `json:"token" form:"token"`
	OrgBin *string   `json:"org_bin" form:"org_bin"`
}

type DeviceCUSt struct {
	Id     *string `json:"id" db:"id"`
	Token  *string `json:"token" db:"token"`
	OrgBin *string `json:"org_bin" db:"org_bin"`
}
