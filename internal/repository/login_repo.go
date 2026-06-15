package repository

func LoginRepo(login string) (id, role int, password string, err error) {
	err = db.QueryRow(
		"SELECT id,pass,role FROM users WHERE login=$1",
		login,
	).Scan(&id, &password, &role)

	return id, role, password, err
}
