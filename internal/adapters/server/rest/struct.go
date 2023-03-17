package rest

type ErrRep struct {
	ErrorCode string `json:"error_code"`
	Desc      string `json:"desc,omitempty"`
}
