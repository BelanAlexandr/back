package repository

func AuthorisedCheck(id int) (rol int, err error) {
	err = db.QueryRow(
		"SELECT role FROM users WHERE id=$1",
		id,
	).Scan(&rol)

	return rol, err
}
