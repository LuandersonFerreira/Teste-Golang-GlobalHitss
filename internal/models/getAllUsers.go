package models

import "github.com/LuandersonFerreira/teste-golang-global/internal/db"

func GetAllUsers() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Uuid, &user.Nome, &user.Sobrenome, &user.Contato, &user.Endereco, &user.Nascimento, &user.Cpf)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
