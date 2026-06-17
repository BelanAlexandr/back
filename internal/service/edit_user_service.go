package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func EditUserService(user models.User) error {
	return repository.EditUserRepo(user)
}
