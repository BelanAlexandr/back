package repository

import "fmt"

func AddUserRepo(login, password string, role int) error {
	var exists bool
	_ = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE login=$1)",
		login,
	).Scan(&exists)

	if exists {
		return fmt.Errorf("Такой логин уже есть")
	}
	_, err := db.Exec(
		"INSERT INTO users(login,pass,role) VALUES($1, $2,$3)",
		login,
		password,
		role,
	)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Ошибка добавления в бд")
	}
	return nil
}
