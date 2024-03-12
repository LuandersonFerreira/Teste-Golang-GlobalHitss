package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
)

func ListUser(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers()
	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for i := range users {
		users[i].Cpf = maskCPF(users[i].Cpf)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
