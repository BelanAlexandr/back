package repository

import (
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

// Предполагается, что db *sql.DB доступен в пакете repository

func EditUserRepo(user models.User) error {
	var query string
	var err error

	if user.Password == "" {
		// Обновление БЕЗ изменения пароля
		query = `
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
		_, err = db.Exec(
			query,
			user.Login,
			user.Role,
			user.Name,
			user.Second_Name,
			user.Middle_Name,
			user.Email,
			user.Phone_Number,
			user.Id, // $8 соответствует user.Id
		)
	} else {
		// Полное обновление ВМЕСТЕ с паролем
		// ВНИМАНИЕ: тут теперь 9 параметров! Исправлены индексы $7, $8, $9
		query = `
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
		_, err = db.Exec(
			query,
			user.Login,
			user.Password,
			user.Role,
			user.Name,
			user.Second_Name,
			user.Middle_Name,
			user.Email,
			user.Phone_Number,
			user.Id, // Теперь это $9, все четко совпадает
		)
	}

	if err != nil {
		fmt.Println("Ошибка при обновлении пользователя в БД:", err)
		return fmt.Errorf("не удалось обновить данные пользователя: %w", err)
	}

	return nil
}
