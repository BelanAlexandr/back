package repository

import "github.com/BelanAlexandr/back/internal/models"

func GetRegionsRepo() (regions []models.Regions, err error) {
	rows, err := db.Query("SELECT te, rus_name FROM dict_region WHERE cd=0 AND ef=0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c models.Regions
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		regions = append(regions, c)
	}
	return regions, nil
}
