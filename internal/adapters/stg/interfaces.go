package stg

type Stg interface {
	CityList(pars *CityListParsSt) ([]*CitySt, error)
}
