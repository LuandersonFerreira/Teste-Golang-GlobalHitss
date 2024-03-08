package models

import "github.com/LuandersonFerreira/teste-golang-global/internal/db"

func UpdateUser(user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET nome=$2, sobrenome=$3, contato=$4, endereco=$5, nascimento=$6, cpf=$7  WHERE uuid=$1`, user.Uuid, user.Nome, user.Sobrenome, user.Contato, user.Endereco, user.Nascimento, user.Cpf)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
