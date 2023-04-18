package entities

type CitySt struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CityListSt struct {
	CitySt
}

type CityListParsSt struct {
	dopTypes.ListParams

	Ids  *[]int64 `json:"ids" form:"ids"`
	Name *string  `json:"name" form:"name"`
}

type CityCUSt struct {
	Name *string `json:"name" db:"name"`
}
