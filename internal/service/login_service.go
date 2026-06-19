package service

import (
	"fmt"

	"github.com/BelanAlexandr/back/internal/repository"
	"github.com/BelanAlexandr/back/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(login, password string) (string, int, error) {

	user_id, user_role, user_password, err := repository.LoginRepo(login)
	if err != nil {
		return "", 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_password), []byte(password))
	if err != nil {
		return "", 0, fmt.Errorf("invalid password")
	}
	token, sessionId, err := utils.GenerateToken(user_id)
	if err != nil {
		return "", 0, err
	}
	err = repository.UpdateUserSession(user_id, sessionId)
	if err != nil {
		return "", 0, err
	}
	return token, user_role, nil
}
