package entities

type SrcSt struct {
	Id        string `json:"id" db:"id"`
	NotifyUrl string `json:"notify_url" db:"notify_url"`
}

type SrcListParsSt struct {
	Ids *[]string `json:"ids" form:"ids"`
}

type SrcCUSt struct {
	Id        *string `json:"id" db:"id"`
	NotifyUrl *string `json:"notify_url" db:"notify_url"`
}
