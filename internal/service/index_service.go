package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(id, role int) ([]models.Exp, error) {
	switch role {
	case models.RoleAdmin:
		return repository.IndexGetRepo()

	case models.RoleDirector:
		return repository.IndexGetRepo()
	case models.RoleEmployee:
		return repository.IndexGetEmployeeRepo(id)
	default:
	}

	return []models.Exp{}, nil
}
