package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(limit, last_id, id_user, role int) ([]models.Exp, error) {
	switch role {
	case models.RoleAdmin:
		return repository.IndexGetRepo(last_id, limit)

	case models.RoleDirector:
		return repository.IndexGetRepo(last_id, limit)
	case models.RoleEmployee:
		return repository.IndexGetEmployeeRepo(id_user, last_id, limit)
	default:
	}

	return []models.Exp{}, nil
}
