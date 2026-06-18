package repository

import (
	"context"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

// Добавили creatorID в параметры, чтобы уведомление шло создателю
func AddUserRepo(creatorID int, user models.User) error {
	ctx := context.Background()

	// Открываем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Ошибка начала транзакции: %w", err)
	}
	// Если что-то пойдет не так, транзакция откатится автоматически
	defer tx.Rollback()

	var exists bool
	// Используем QueryRowContext внутри транзакции
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
		user.Name,         // сопоставляется с first_name
		user.Second_Name,  // сопоставляется с last_name
		user.Middle_Name,  // сопоставляется с middle_name
		user.Email,        // сопоставляется с email
		user.Phone_Number, // сопоставляется с phone
	)

	if err != nil {
		fmt.Println("Критическая ошибка БД при вставке:", err)
		return fmt.Errorf("Ошибка добавления в бд: %w", err)
	}

	// Отправляем уведомление администратору (creatorID) о создании нового сотрудника
	message := fmt.Sprintf("Успешно добавлен новый пользователь с логином: %s", user.Login)
	_, err = AddNotification(ctx, tx, creatorID, message)
	if err != nil {
		return fmt.Errorf("Ошибка отправки уведомления: %w", err)
	}

	// Фиксируем транзакцию
	return tx.Commit()
}
