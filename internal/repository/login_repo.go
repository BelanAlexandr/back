package repository

func LoginRepo(login string) (id int, password string, err error) {
	err = db.QueryRow(
		"SELECT id,pass FROM users WHERE login=$1",
		login,
	).Scan(&id, &password)

	return id, password, err
}
