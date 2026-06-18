package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

// Предполагается, что db *sql.DB доступен в пакете repository

// Добавили creatorID, чтобы отправлять уведомление администратору
func EditUserRepo(creatorID int, user models.User) error {
	ctx := context.Background()

	// 1. Открываем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("ошибка начала транзакции: %w", err)
	}
	// Автоматический откат при ошибке или панике
	defer tx.Rollback()

	var result sql.Result

	if user.Password == "" {
		// Обновление БЕЗ изменения пароля
		query := `
			UPDATE users 
			SET login = $1, 
				role = $2, 
				first_name = $3, 
				last_name = $4, 
				middle_name = $5, 
				email = NULLIF($6, ''), 
				phone = NULLIF($7, '')
			WHERE id = $8
		`
		result, err = tx.ExecContext(
			ctx,
			query,
			user.Login,
			user.Role,
			user.Name,
			user.Second_Name,
			user.Middle_Name,
			user.Email,
			user.Phone_Number,
			user.Id,
		)
	} else {
		// Полное обновление ВМЕСТЕ с паролем
		query := `
			UPDATE users 
			SET login = $1, 
				pass = $2, 
				role = $3, 
				first_name = $4, 
				last_name = $5, 
				middle_name = $6, 
				email = NULLIF($7, ''), 
				phone = NULLIF($8, '')
			WHERE id = $9
		`
		result, err = tx.ExecContext(
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
			user.Id,
		)
	}

	if err != nil {
		fmt.Println("Ошибка при обновлении пользователя в БД:", err)
		return fmt.Errorf("не удалось обновить данные пользователя: %w", err)
	}

	// Проверяем, обновился ли хоть один пользователь
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("пользователь с id %d не найден для обновления", user.Id)
	}

	// 2. ВЫЗОВ УВЕДОМЛЕНИЯ ВНУТРИ ТРАНЗАКЦИИ
	message := fmt.Sprintf("Обновлены данные пользователя с логином: %s (ID: %d)", user.Login, user.Id)
	_, err = AddNotification(ctx, tx, creatorID, message)
	if err != nil {
		return fmt.Errorf("ошибка отправки уведомления: %w", err)
	}

	// 3. Коммитим транзакцию
	return tx.Commit()
}
