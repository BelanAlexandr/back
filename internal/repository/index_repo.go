package repository

import (
	"github.com/BelanAlexandr/back/internal/models"
)

func IndexGetRepo() ([]models.FilesNames, error) {
	query := `
    SELECT* (file_name,is_closed) FROM filenames  
   `

	res, err := db.Query(query)
	if err != nil {
		return []models.FilesNames{}, err
	}
	var files []models.FilesNames
	for res.Next() {
		var file models.FilesNames
		err := res.Scan(&file.File_Name, file.Is_closed)
		if err != nil {
			return []models.FilesNames{}, err
		}
		files = append(files, file)
	}
	return files, nil
}
