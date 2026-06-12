package service

import (
	"slices"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func GetRegionsService() (reg []models.Regions, err error) {
	reg, err = repository.GetRegionsRepo()
	if err != nil {
		return nil, err
	}
	slices.SortFunc(reg, func(a, b models.Regions) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})
	return reg, nil
}
