package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(id, role int) []models.FilesNames {
	switch role {
	case models.RoleAdmin:
		repository.IndexGetRepo()
	case models.RoleDirector:
	case models.RoleEmployee:
	default:
	}

	return []models.FilesNames{}
}
