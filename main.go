package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}
