package repository

import (
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddUserRepo(user models.User) error {
	var exists bool

	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		user.Login,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Ошибка проверки логина: %w", err)
	}

	if exists {
		return fmt.Errorf("Такой логин уже есть")
	}

	// Используем NULLIF($7, '') и NULLIF($8, ''), чтобы пустые строки превращались в NULL
	query := `
        INSERT INTO users (login, pass, role, first_name, last_name, middle_name, email, phone) 
        VALUES ($1, $2, $3, $4, $5, $6, NULLIF($7, ''), NULLIF($8, ''))
    `

	_, err = db.Exec(
		query,
		user.Login,
		user.Password,
		user.Role,
		user.Name,         // сопоставляется с first_name
		user.Second_Name,  // сопоставляется с last_name
		user.Middle_Name,  // сопоставляется с middle_name
		user.Email,        // сопоставляется с email
		user.Phone_Number, // сопоставляется с phone
	)

	if err != nil {
		fmt.Println("Критическая ошибка БД при вставке:", err)
		return fmt.Errorf("Ошибка добавления в бд")
	}

	return nil
}
