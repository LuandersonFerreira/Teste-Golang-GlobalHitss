package models

import "github.com/LuandersonFerreira/teste-golang-global/internal/db"

func DeleteUser(uuid string) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE uuid = $1`, uuid)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
