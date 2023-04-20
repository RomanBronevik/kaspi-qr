package provider

// request

type TradePointCreateReqSt struct {
	Name            string `json:"name"`
	TradePointId    string `json:"trade_point_id"`
	OrganizationBin string `json:"organization_id"`
}

// reply

type TradePointListRepSt struct {
	BaseRepSt
	Data []*TradePointSt `json:"Data"`
}

// common

type TradePointSt struct {
	TradePointId   int64  `json:"TradePointId"`
	TradePointName string `json:"TradePointName"`
}
