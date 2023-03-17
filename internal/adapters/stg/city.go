package stg

type CitySt struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	SiteCode string `json:"site_code"`
}

type CityListParsSt struct {
	HasSiteCode *bool `json:"has_site_code" form:"has_site_code"`
}
