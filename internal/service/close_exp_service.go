package service

import (
	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func CloseExpService(exp models.Exp) error {
	getFloat := func(f *float64) float64 {
		if f == nil {
			return 0.0
		}
		return *f
	}

	total := getFloat(exp.Expert_Cost) + getFloat(exp.Material_Cost) + getFloat(exp.Exploitation_Cost)
	exp.Full_Cost = &total
	total = total + total*0.16
	exp.Full_Cost_Nds = &total
	return repository.CloseExpRepo(exp, true)
}
