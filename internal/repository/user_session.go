package repository

func UpdateUserSession(user_id int, session_id string) (err error) {
	_, err = db.Exec("UPDATE users SET current_session_id=$1 WHERE id=$2 ", session_id, user_id)

	return err
}
func GetUserSessionID(user_id int) (string, error) {
	var sessionID string
	err := db.QueryRow("SELECT current_session_id FROM users WHERE id = $1", user_id).Scan(&sessionID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}
