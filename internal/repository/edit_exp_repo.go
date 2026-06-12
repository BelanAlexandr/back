package repository

import "github.com/BelanAlexandr/back/internal/models"

func EditExpRepo(exp models.Exp, closed bool) error {
	return UpdateExpRepo(exp, false)
}
