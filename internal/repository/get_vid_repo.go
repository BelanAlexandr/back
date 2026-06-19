package repository

import "github.com/BelanAlexandr/back/internal/models"

func GetVidRepo() (vid []models.Vid, err error) {
	rows, err := db.Query("SELECT name, shifr FROM dict_vid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var c models.Vid
		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		vid = append(vid, c)
	}
	return vid, nil
}
