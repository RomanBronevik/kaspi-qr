package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	_ "kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/domain/entities"
)

func (r *St) CreateCity(ctx context.Context, city *entities.CreateCityDTO) error {
	q := `
		INSERT INTO city (name, organization_bin, code) 
		VALUES ($1, $2, $3)`

	if err := r.db.Exec(ctx, q, city.Name, city.OrganizationBin, city.Code); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) FindAllCities(ctx context.Context) (u []entities.City, err error) {
	q := `
		SELECT code, name,  organization_bin FROM city`
	rows, err := r.db.Query(ctx, q)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	defer rows.Close()

	cities := make([]entities.City, 0)

	for rows.Next() {
		var city entities.City

		err := rows.Scan(&city.Code, &city.Name, &city.OrganizationBin)
		if err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}

func (r *St) FindOneCityByCityCode(ctx context.Context, code string) (entities.City, error) {
	q := `
		SELECT code, name,  organization_bin FROM city WHERE code = $1`

	//Trace
	var city entities.City
	err := r.db.QueryRow(ctx, q, code).Scan(&city.Code, &city.Name, &city.OrganizationBin)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return entities.City{}, err
	}

	return city, nil
}

func (r *St) DeleteCity(ctx context.Context, id string) error {
	q := `
		DELETE FROM city
		WHERE id = $1;`

	if err := r.db.Exec(ctx, q, id); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}

func (r *St) DeleteCities(ctx context.Context) error {
	q := `
		TRUNCATE TABLE city;`

	if err := r.db.Exec(ctx, q); err != nil {
		return r.ErorrHandler(err)
	}

	return nil
}
