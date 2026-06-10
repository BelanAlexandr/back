package service

import (
	"fmt"

	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/BelanAlexandr/back/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(login, password string) (string, error) {

	user_id, user_password, err := repository.LoginRepo(login)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid password")
	}
	token, err := utils.GenerateToken(user_id)
	if err != nil {
		return "", err
	}
	return token, nil
}
