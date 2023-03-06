package entities

type Organization struct {
	Name string `json:"name"`
	Bin  string `json:"bin"`
}

type CreateOrganizationDTO struct {
	Name string `json:"name"`
	Bin  string `json:"bin"`
}
