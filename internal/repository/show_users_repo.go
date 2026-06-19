package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/BelanAlexandr/back/internal/models"
)

func ShowUsersRepo(offset, limit int, sortField, sortOrder string, searchQuery, roleFilter string) ([]models.User, int, error) {
	var users []models.User
	var totalCount int

	selectQuery := "SELECT id, login, role, first_name, last_name,middle_name, email, phone FROM users"
	countQuery := "SELECT COUNT(*) FROM users"

	var conditions []string
	var args []interface{}
	argCounter := 1

	if searchQuery != "" {

		searchParam := "%" + strings.ToLower(searchQuery) + "%"
		conditions = append(conditions, fmt.Sprintf(
			"(LOWER(login) LIKE $%d OR LOWER(email) LIKE $%d OR LOWER(first_name) LIKE $%d OR LOWER(last_name) LIKE $%d)",
			argCounter, argCounter, argCounter, argCounter,
		))
		args = append(args, searchParam)
		argCounter++
	}

	if roleFilter != "" {
		roleID, err := strconv.Atoi(roleFilter)
		if err == nil {
			conditions = append(conditions, fmt.Sprintf("role = $%d", argCounter))
			args = append(args, roleID)
			argCounter++
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	fullCountQuery := countQuery + whereClause
	err := db.QueryRow(fullCountQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("ошибка подсчета количества пользователей: %w", err)
	}

	if totalCount == 0 {
		return []models.User{}, 0, nil
	}

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

		var firstName, lastName, middlename, email, phone sql.NullString

		err := rows.Scan(&u.Id, &u.Login, &u.Role, &firstName, &lastName, &middlename, &email, &phone)
		if err != nil {
			return nil, 0, fmt.Errorf("ошибка сканирования пользователя: %w", err)
		}

		u.Name = firstName.String
		u.Second_Name = lastName.String
		u.Middle_Name = middlename.String
		u.Email = email.String
		u.Phone_Number = phone.String
		u.Password = ""
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
