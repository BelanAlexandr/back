package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func EditUserService(creator_id int, user models.User) error {
	return repository.EditUserRepo(creator_id, user)
}
