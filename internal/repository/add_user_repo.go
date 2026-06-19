package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddUserRepo(creatorID int, user models.User) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Ошибка начала транзакции: %w", err)
	}

	defer tx.Rollback()

	var exists bool

	err = tx.QueryRowContext(
		ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		user.Login,
	).Scan(&exists)
	if err != nil {
		return fmt.Errorf("Ошибка проверки логина: %w", err)
	}

	if exists {
		return fmt.Errorf("Такой логин уже есть")
	}

	query := `
		INSERT INTO users (login, pass, role, first_name, last_name, middle_name, email, phone) 
		VALUES ($1, $2, $3, $4, $5, $6, NULLIF($7, ''), NULLIF($8, ''))
	`

	_, err = tx.ExecContext(
		ctx,
		query,
		user.Login,
		user.Password,
		user.Role,
		user.Name,
		user.Second_Name,
		user.Middle_Name,
		user.Email,
		user.Phone_Number,
	)

	if err != nil {
		fmt.Println("Критическая ошибка БД при вставке:", err)
		return fmt.Errorf("Ошибка добавления в бд: %w", err)
	}

	message := fmt.Sprintf("Успешно добавлен новый пользователь с логином: %s", user.Login)
	_, err = AddNotification(ctx, tx, creatorID, message)
	if err != nil {
		return fmt.Errorf("Ошибка отправки уведомления: %w", err)
	}

	return tx.Commit()
}
