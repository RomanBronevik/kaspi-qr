package rest

import (
	"fmt"
	"net/url"

	"kaspi-qr/internal/adapters/stg"
)

func (o *St) CityList(pars *stg.CityListParsSt) ([]*stg.CitySt, error) {
	qPars := url.Values{}

	if pars.HasSiteCode != nil {
		qPars.Set("has_site_code", fmt.Sprintf("%v", *pars.HasSiteCode))
	}

	result := make([]*stg.CitySt, 0)

	_, err := o.sendRequest("GET", "city", nil, qPars, &result)
	if err != nil {
		o.lg.Errorw("Stg: CityList", err)
		return nil, err
	}

	return result, nil
}
