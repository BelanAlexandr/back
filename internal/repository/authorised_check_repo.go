package repository

func AuthorisedCheck(id int) (rol int, session string, err error) {
	err = db.QueryRow(
		"SELECT role,current_session_id FROM users WHERE id=$1",
		id,
	).Scan(&rol, &session)

	return rol, session, err
}
