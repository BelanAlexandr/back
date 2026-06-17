package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func ShowUsersService(offset, limit int,
	sortField, sortOrder string,
	searchQuery, roleFilter string) ([]models.User, int, error) {
	return repository.ShowUsersRepo(offset, limit, sortField, sortOrder, searchQuery, roleFilter)
}
