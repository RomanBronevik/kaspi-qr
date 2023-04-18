package entities

type OrganisationSt struct {
	Id   string `json:"id" db:"id"`
	Bin  string `json:"bin" db:"bin"`
	Name string `json:"name" db:"name"`
}

type OrganisationListParsSt struct {
	Ids  *[]int64 `json:"ids" form:"ids"`
	Name *string  `json:"name" form:"name"`
}

type OrganisationCUSt struct {
	Id   *string `json:"id" db:"id"`
	Bin  *string `json:"bin" db:"bin"`
	Name *string `json:"name" db:"name"`
}
