package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
	"github.com/LuandersonFerreira/teste-golang-global/internal/rabbitmq"
	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user.Uuid = uuid.New()

	err = rabbitmq.ConnectQueue(user)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir um usuário %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Usuário adicionado a fila com sucesso! ID: %s", user.Uuid),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
