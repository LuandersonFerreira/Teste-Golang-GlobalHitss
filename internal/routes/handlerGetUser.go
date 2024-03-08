package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LuandersonFerreira/teste-golang-global/internal/models"
	"github.com/google/uuid"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	param := queryParams.Get("uuid")
	if param == "" {
		log.Printf("É necessário enviar o campo uuid")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uuid, err := uuid.Parse(param)
	if err != nil {
		log.Printf("É necessário ser um uuid válido")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := models.GetUser(uuid)
	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
