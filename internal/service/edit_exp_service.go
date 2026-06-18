package service

import (
	"log"

	"github.com/BelanAlexandr/back/internal/models"
	"github.com/BelanAlexandr/back/internal/repository"
)

func EditExpService(exp models.Exp) error {

	getFloat := func(f *float64) float64 {
		if f == nil {
			return 0.0
		}
		return *f
	}

	total := getFloat(exp.Expert_Cost) + getFloat(exp.Material_Cost) + getFloat(exp.Exploitation_Cost)
	exp.Full_Cost = &total
	log.Println(total)
	total = total + total*0.16
	exp.Full_Cost_Nds = &total
	log.Println(total)
	return repository.EditExpRepo(exp, false)

}
