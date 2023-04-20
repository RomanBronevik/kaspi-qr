package entities

type TradePointSt struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type TradePointListParsSt struct {
	OrgBin *string `json:"org_bin" form:"org_bin"`
}
