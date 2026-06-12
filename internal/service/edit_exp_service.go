package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func EditExpService(exp models.Exp) error {
	return repository.EditExpRepo(exp, false)
}
