package main

import ( //"fmt"
	"net/http"
)

func main() {

	// persona1 := Persona{
	//     Nombre:      "Juan",
	//     Apellido:    "Perez",
	//     Edad:        30,
	//     CountryCode: "ES",
	// }

	// persona2 := Persona{
	//     Nombre:      "Maria",
	//     Apellido:    "Gomez",
	//     Edad:        25,
	//     CountryCode: "MX",
	// }

	// persona3 := Persona{
	//     Nombre:      "Pedro",
	//     Apellido:    "Martinez",
	//     Edad:        40,
	//     CountryCode: "AR",
	// }

	router := http.NewServeMux()

	router.HandleFunc("/",getUno)

	router.HandleFunc("POST /personas", createPersona)

	router.HandleFunc("PUT /personas", updatePersona)

	router.HandleFunc("GET /personas/{id}",getPerson)

	router.HandleFunc("GET /personas",getPersonas)

	http.ListenAndServe(":8000", router)

}

func getUno(w http.ResponseWriter,router * http.Request){
	w.Write([]byte("Mi primera api"))
}