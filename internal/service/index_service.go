package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(id, role int) ([]models.FilesNames, error) {
	switch role {
	case models.RoleAdmin:
		return repository.IndexGetRepo()

	case models.RoleDirector:
	case models.RoleEmployee:
	default:
	}

	return []models.FilesNames{}, nil
}
