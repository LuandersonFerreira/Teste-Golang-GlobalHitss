package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	uuid := queryParams.Get("uuid")
	if uuid == "" {
		log.Printf("É necessário enviar o campo uuid")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any

	_, err := models.DeleteUser(uuid)
	if err != nil {
		log.Printf("Erro ao fazer update do usuário: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Usuário deletado com sucesso! ID: %s", uuid),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
