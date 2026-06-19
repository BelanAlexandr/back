package service

import (
	"errors"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func IndexGetService(id_user, role, offset, limit int, sortField, sortOrder, statusFilter, dateFrom, dateTo string) ([]models.ExpListItem, int, error) {
	switch role {
	case models.RoleAdmin, models.RoleDirector:

		return repository.IndexGetRepo(offset, limit, sortField, sortOrder, statusFilter, dateFrom, dateTo)

	case models.RoleEmployee:

		return repository.IndexGetEmployeeRepo(id_user, offset, limit, sortField, sortOrder, statusFilter, dateFrom, dateTo)

	default:
		return nil, 0, errors.New("неизвестная роль пользователя")
	}
}
