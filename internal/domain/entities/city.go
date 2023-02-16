package entities

type CreateCityDTO struct {
	Name            string `json:"name"`
	OrganizationBin string `json:"organization_bin"`
}

type City struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	OrganizationBin string `json:"organization_bin"`
}
