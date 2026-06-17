package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/BelanAlexandr/back/internal/models"
)

// db — ваш глобальный или локальный пул соединений *sql.DB

func ShowUsersRepo(offset, limit int, sortField, sortOrder string, searchQuery, roleFilter string) ([]models.User, int, error) {
	var users []models.User
	var totalCount int

	// Базовые запросы
	selectQuery := "SELECT id, login, pass, role, first_name, last_name, email, phone FROM users"
	countQuery := "SELECT COUNT(*) FROM users"

	var conditions []string
	var args []interface{}
	argCounter := 1

	// 1. Фильтр по поисковой строке (ищем совпадения в login, email, first_name, last_name)
	if searchQuery != "" {
		// Приводим к нижнему регистру для регистронезависимого поиска
		searchParam := "%" + strings.ToLower(searchQuery) + "%"
		conditions = append(conditions, fmt.Sprintf(
			"(LOWER(login) LIKE $%d OR LOWER(email) LIKE $%d OR LOWER(first_name) LIKE $%d OR LOWER(last_name) LIKE $%d)",
			argCounter, argCounter, argCounter, argCounter,
		))
		args = append(args, searchParam)
		argCounter++
	}

	// 2. Фильтр по роли
	if roleFilter != "" {
		roleID, err := strconv.Atoi(roleFilter)
		if err == nil {
			conditions = append(conditions, fmt.Sprintf("role = $%d", argCounter))
			args = append(args, roleID)
			argCounter++
		}
	}

	// Объединяем условия WHERE, если они есть
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// --- СНАЧАЛА СЧИТАЕМ TOTAL COUNT ---
	fullCountQuery := countQuery + whereClause
	err := db.QueryRow(fullCountQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка подсчета количества пользователей: %w", err)
	}

	// Если записей вообще нет, можно не делать второй запрос
	if totalCount == 0 {
		return []models.User{}, 0, nil
	}

	// --- ТЕПЕРЬ ПОЛУЧАЕМ СТРОКИ С СОРТИРОВКОЙ И ПАГИНАЦИЕЙ ---
	// Безопасно подставляем sortField и sortOrder, так как они проверены в хендлере (white-list)
	fullSelectQuery := fmt.Sprintf("%s%s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		selectQuery, whereClause, sortField, sortOrder, argCounter, argCounter+1,
	)
	args = append(args, limit, offset)

	rows, err := db.Query(fullSelectQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка выполнения запроса пользователей: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		// Используем sql.NullString на случай, если в БД новые поля пока еще NULL
		var firstName, lastName, email, phone sql.NullString

		err := rows.Scan(&u.Id, &u.Login, &u.Password, &u.Role, &firstName, &lastName, &email, &phone)
		if err != nil {
			return nil, 0, fmt.Errorf("ошибка сканирования пользователя: %w", err)
		}

		// Переносим значения из NullString в обычные строки нашей структуры
		u.Name = firstName.String
		u.Second_Name = lastName.String
		u.Email = email.String
		u.Phone_Number = phone.String

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
