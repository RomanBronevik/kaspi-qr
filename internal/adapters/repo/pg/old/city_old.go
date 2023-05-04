package old

import (
	"context"
	"errors"
	_ "kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/domain/entities"

	"github.com/rendau/dop/dopErrs"
)

func (d *St) CreateCity(ctx context.Context, city *entities.CreateCityDTO) error {
	q := `
		INSERT INTO city (name, organization_bin, code) 
		VALUES ($1, $2, $3)`

	return d.DbExec(ctx, q, city.Name, city.OrganizationBin, city.Code)
}

func (d *St) FindAllCities(ctx context.Context) (u []*entities.City, err error) {
	q := `
		SELECT code, name,  organization_bin
		FROM city
		order by name`

	rows, err := d.DbQuery(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := make([]*entities.City, 0)

	for rows.Next() {
		city := &entities.City{}

		err = rows.Scan(&city.Code, &city.Name, &city.OrganizationBin)
		if err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func (d *St) FindOneCityByCityCode(ctx context.Context, code string) (*entities.City, error) {
	q := `
		SELECT code, name,  organization_bin
		FROM city
		WHERE code = $1`

	city := &entities.City{}

	err := d.DbQueryRow(ctx, q, code).Scan(&city.Code, &city.Name, &city.OrganizationBin)
	if err != nil {
		if errors.Is(err, dopErrs.NoRows) {
			return nil, nil
		}
		return nil, err
	}

	return city, nil
}

func (d *St) DeleteCity(ctx context.Context, id string) error {
	q := `
		DELETE FROM city
		WHERE id = $1;`

	return d.DbExec(ctx, q, id)
}

func (d *St) DeleteCities(ctx context.Context) error {
	q := `
		TRUNCATE TABLE city;`

	return d.DbExec(ctx, q)
}
