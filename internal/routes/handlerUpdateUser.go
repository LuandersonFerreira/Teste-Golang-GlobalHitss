package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any

	_, err = models.UpdateUser(user)
	if err != nil {
		log.Printf("Erro ao fazer update do usuário: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Usuário atualizado com sucesso! ID: %s", user.Uuid),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
