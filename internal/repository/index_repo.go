package repository

import (
	"github.com/BelanAlexandr/back/internal/models"
)

func IndexGetRepo() ([]models.Exp, error) {
	query := `
    SELECT id, is_closed 
    FROM electronic_journal;
`
	res, err := db.Query(query)
	if err != nil {
		return []models.Exp{}, err
	}

	var files []models.Exp
	for res.Next() {
		var file models.Exp
		err := res.Scan(&file.Id, &file.Is_Closed)
		if err != nil {
			return []models.Exp{}, err
		}

		files = append(files, file)
	}

	return files, nil
}
func IndexGetEmployeeRepo(id int) ([]models.Exp, error) {
	query := `
    SELECT id, is_closed 
    FROM electronic_journal
	WHERE creator_id=$1;
`
	res, err := db.Query(query, id)
	if err != nil {
		return []models.Exp{}, err
	}

	var files []models.Exp
	for res.Next() {
		var file models.Exp
		err := res.Scan(&file.Id, &file.Is_Closed)
		if err != nil {
			return []models.Exp{}, err
		}

		files = append(files, file)
	}

	return files, nil
}
