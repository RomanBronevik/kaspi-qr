package entities

type ReturnSt struct {
	StatusCode            int                    `json:"StatusCode"`
	Message               string                 `json:"Message"`
	ReturnOperationDataSt *ReturnOperationDataSt `json:"ReturnOperationDataSt"`
}

type ReturnOperationDataSt struct {
	ReturnOperationId int `json:"ReturnOperationId"`
}
