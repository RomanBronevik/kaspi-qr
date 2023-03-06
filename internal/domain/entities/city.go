package entities

type CreateCityDTO struct {
	Name            string `json:"name"`
	OrganizationBin string `json:"organization_bin"`
	Code            string `json:"code"`
}

type City struct {
	Name            string `json:"name"`
	OrganizationBin string `json:"organization_bin"`
	Code            string `json:"code"`
}

type CityUpdateReqOutput struct {
	Result bool `json:"result"`
	//Errors *Errors     `json:"errors"`
	Data *Cities `json:"data"`
}

type Cities struct {
	Cities []CitiesData `json:"cities"`
}

type CitiesData struct {
	Code        string       `json:"code"`
	Name        string       `json:"name"`
	Phones      []string     `json:"phones"`
	Coordinates *Coordinates `json:"coordinates"`
	Translit    string       `json:"translit"`
}

type Errors struct {
}

type Coordinates struct {
	Lan float64 `json:"lan"`
	Lon float64 `json:"lon"`
}
