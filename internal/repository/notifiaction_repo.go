package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/BelanAlexandr/back/internal/models"
)

func AddNotification(ctx context.Context, tx *sql.Tx, userID int, message string) (int, error) {
	querySave := `
    INSERT INTO notifications (user_id, mess, is_read, created_at) 
    VALUES ($1, $2, false, NOW())
    RETURNING id`

	var lastInsertId int
	var err error

	// Если транзакция передана, выполняем в ней. Если нет — используем стандартный db
	if tx != nil {
		err = tx.QueryRowContext(ctx, querySave, userID, message).Scan(&lastInsertId)
	} else {
		err = db.QueryRowContext(ctx, querySave, userID, message).Scan(&lastInsertId)
	}

	if err != nil {
		log.Println("Ошибка сохранения уведомления в БД:", err)
		return 0, err
	}

	return lastInsertId, nil
}
func MarkNotification(message_id, user_id int) error {
	query := "UPDATE notifications SET is_read = true WHERE id = $1 AND user_id = $2"
	_, err := db.Exec(query, message_id, user_id)
	return err
}
func GetNotifications(user_id int) ([]models.Notification, error) {
	queryList := `
    SELECT id, user_id, text, is_read, created_at 
    FROM notifications 
    WHERE user_id = $1
    ORDER BY id DESC 
    LIMIT 20`

	rows, err := db.Query(queryList, user_id)
	if err != nil {

		return []models.Notification{}, err
	}
	defer rows.Close()

	var notificationsFromDB []models.Notification
	for rows.Next() {
		var n models.Notification
		err := rows.Scan(&n.ID, &n.User_ID, &n.Text, &n.Is_read, &n.Created_at)
		if err != nil {

			return []models.Notification{}, err
		}
		notificationsFromDB = append(notificationsFromDB, n)
	}
	return notificationsFromDB, nil
}
