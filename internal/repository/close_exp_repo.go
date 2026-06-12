package repository

import (
	"github.com/BelanAlexandr/back/internal/models"
)

func CloseExpRepo(exp models.Exp, closed bool) error {
	return UpdateExpRepo(exp, true)
}
