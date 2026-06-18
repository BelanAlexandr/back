package repository

import (
	"log"
)

func AddNotification(user_id int, message string) (int, error) {
	querySave := `
    INSERT INTO notifications (user_id, text, is_read, created_at) 
    VALUES ($1, $2, false, NOW())
    RETURNING id`

	var lastInsertId int

	err := db.QueryRow(querySave, user_id, message).Scan(&lastInsertId)
	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
		return 0, err
	}

	return lastInsertId, nil
}
