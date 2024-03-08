package models

import (
	"time"

	"github.com/LuandersonFerreira/teste-golang-global/internal/db"
)

const customDateFormat = "02/01/2006"

func InsertUser(user User) (uuid string, err error) {

	parsedTime, err := time.Parse(customDateFormat, user.Nascimento)

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users (uuid, nome, sobrenome, contato, endereco, nascimento, cpf) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING uuid`

	err = conn.QueryRow(sql, user.Uuid, user.Nome, user.Sobrenome, user.Contato, user.Endereco, parsedTime, user.Cpf).Scan(&uuid)

	return
}
