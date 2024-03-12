package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	maskedcpf := maskCPF(user.Cpf)
	user.Cpf = maskedcpf

	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func maskCPF(cpf string) string {
	cpfNumbers := strings.ReplaceAll(cpf, ".", "")
	cpfNumbers = strings.ReplaceAll(cpfNumbers, "-", "")

	cpfPrefix := cpfNumbers[:3]

	maskedCpf := cpfPrefix + strings.Repeat("*", len(cpfNumbers)-3)

	maskedCpf = fmt.Sprintf("%s.%s.%s-%s", maskedCpf[:3], maskedCpf[3:6], maskedCpf[6:9], maskedCpf[9:])

	return maskedCpf
}
