package main

import (
	"log"
	"net/http"
	"rinha_de_devs___golang/modules"

	"github.com/gorilla/mux"
)

func main() {
	r := modules.NewRepositoryTransaction()
	s := modules.NewService(r)
	handler := modules.NewHandler(s)

	router := mux.NewRouter()
	router.HandleFunc("/ping", modules.Ping).Methods("GET")
	router.HandleFunc("/clientes/{id}/transacoes", handler.PostTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
