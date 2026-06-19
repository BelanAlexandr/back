package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func GetVidService() (vid []models.Vid, err error) {
	vid, err = repository.GetVidRepo()
	return vid, nil
}
