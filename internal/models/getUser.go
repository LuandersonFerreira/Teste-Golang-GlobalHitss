package models

import (
	"github.com/LuandersonFerreira/teste-golang-global/internal/db"
	"github.com/google/uuid"
)

func GetUser(uuid uuid.UUID) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE uuid=$1`, uuid)

	err = row.Scan(&user.Uuid, &user.Nome, &user.Sobrenome, &user.Contato, &user.Endereco, &user.Nascimento, &user.Cpf)

	return
}
