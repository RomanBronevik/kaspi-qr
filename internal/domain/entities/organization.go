package entities

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Bin  string `json:"bin"`
}

type CreateOrganizationDTO struct {
	Name string `json:"name"`
	Bin  string `json:"bin"`
}
