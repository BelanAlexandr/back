package service

import (
	"github.com/BelanAlexandr/back/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func AddUserService(login, password string, role int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return repository.AddUserRepo(login, string(hashedPassword), role)

}
