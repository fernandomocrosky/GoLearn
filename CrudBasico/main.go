package main

import (
	"crud/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.UpdateUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeleteUsuario).Methods(http.MethodDelete)

	println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
