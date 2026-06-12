package service

import (
	"errors"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(id_user, role, last_id, limit int, statusFilter, dateFrom, dateTo string) ([]models.Exp, error) {
	switch role {
	case models.RoleAdmin, models.RoleDirector:

		return repository.IndexGetRepo(last_id, limit, statusFilter, dateFrom, dateTo)

	case models.RoleEmployee:

		return repository.IndexGetEmployeeRepo(id_user, last_id, limit, statusFilter, dateFrom, dateTo)

	default:

		return nil, errors.New("неизвестная роль пользователя")
	}
}
