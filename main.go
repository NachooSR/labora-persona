package main

import (
	_ "database/sql"
	//"fmt"

	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	initDb()

	router := http.NewServeMux()

	router.HandleFunc("/", getUno)

	router.HandleFunc("POST /personas", createPersona)

	router.HandleFunc("PUT /personas", updatePersona)

	router.HandleFunc("GET /personas/{id}", getPerson)

	router.HandleFunc("DELETE/personas/{id}", DeletePerson)

	http.ListenAndServe(":8000", router)

}

func getUno(w http.ResponseWriter, router *http.Request) {
	w.Write([]byte("Mi primera api"))
}
