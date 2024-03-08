package models

import (
	"github.com/google/uuid"
)

type User struct {
	Uuid       uuid.UUID
	Nome       string
	Sobrenome  string
	Contato    string
	Endereco   string
	Nascimento string
	Cpf        string
}
