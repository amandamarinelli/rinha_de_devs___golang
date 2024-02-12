package modules

import (
	"encoding/json"
	"net/http"
	"rinha_de_devs___golang/models"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	Handler struct {
		Service Service
	}

	Service interface {
		PostTransaction(id uint64) (models.TransactionResponse, error)
	}
)

func NewHandler(s Service) Handler {
	return Handler{
		Service: s,
	}
}

func (h Handler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		err := "Id é obrigatório"
		json.NewEncoder(w).Encode(err)
		return
	}

	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		err := "Id deve ser um número inteiro"
		json.NewEncoder(w).Encode(err)
		return
	}

	resp, err := h.Service.PostTransaction(userId)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(resp)
}
