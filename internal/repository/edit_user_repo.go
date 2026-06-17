package repository

import (
	"fmt"

	"github.com/BelanAlexandr/back/internal/models"
)

// Предполагается, что db *sql.DB доступен в пакете repository

func EditUserRepo(user models.User) error {
	// 1. Проверяем, передан ли новый пароль.
	// Если пароль пустой, мы обновляем все поля, кроме пароля.
	// Если пароль заполнен, мы обновляем и его тоже (не забудьте захэшировать его перед отправкой в репозиторий!).

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
			user.Id, // ID идет последним аргументом ($8)
		)
	} else {
		// Полное обновление ВМЕСТЕ с паролем
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
			user.Id, // ID идет последним аргументом ($9)
		)
	}

	if err != nil {
		// Печатаем в консоль для отладки, но наружу отдаем понятную ошибку
		fmt.Println("Ошибка при обновлении пользователя в БД:", err)
		return fmt.Errorf("не удалось обновить данные пользователя: %w", err)
	}

	return nil
}
