package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func createPersona(w http.ResponseWriter, r *http.Request) {

	///Cuerpo de la socilitud para bajarlo a una variable
	decoder := json.NewDecoder(r.Body)
	var aux Persona
	err := decoder.Decode(&aux)

	//Error por si no lo pudo decodificar
	if err != nil {
		println(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	if !aux.Validate() {
		w.WriteHeader(http.StatusBadRequest)

	}

	///La persona es ingresada en la bdd, y el id lo usamos para poder enviarlo en la response
	id, errorcito := insertarPersonaEnLaDb(aux)

	if errorcito != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	aux.ID = id
	encoder := json.NewEncoder(w)
	encoder.Encode(aux)
}

func getPerson(w http.ResponseWriter, r *http.Request) {

	idPersona := r.PathValue("id")

	idComoInt, _ := strconv.Atoi(idPersona)

	persona, err := ObtenerPersona(idComoInt)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(persona)
}

func updatePersona(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var personaAux Persona

	err := decoder.Decode(&personaAux)
	if err != nil {
		fmt.Println(w, "ERROR: "+err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	personaNueva, errorcito := editarPersonaEnLaDb(personaAux.ID, personaAux.Edad, personaAux.Nombre, personaAux.Apellido, personaAux.CountryCode)

	if errorcito != nil {
		fmt.Println("ERROR: " + errorcito.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(personaNueva)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

	idPersona := r.PathValue("id")

	idComoInt, _ := strconv.Atoi(idPersona)

	err := eliminarPersona(idComoInt)
	if err != nil {

		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

// Listar todas las personas de La bdd
// func getPersonas(w http.ResponseWriter, _ *http.Request) { ////POCO EFICIENTE
// 	// Aquí implementa la lógica para para obtener todas las personas
// 	mapaInfoCountry:=make(map[string]CountryInfo)
// 	var personasExtendidas[]PersonaExt
// 	var countryInfo CountryInfo
//     var errorcito error

// 	for i := 0; i < len(PersonasDB); i++ {

// 		_,ok:=mapaInfoCountry[PersonasDB[i].CountryCode]
// 		///Si la llave (osea el country code No existe) llamamos a la api

// 		if !ok {
// 		    countryInfo, errorcito = getInfoCountry(PersonasDB[i].CountryCode)
// 		    if errorcito != nil {
// 				http.Error(w, "Error al obtener información del país", http.StatusInternalServerError)
// 				return
// 			}
// 			mapaInfoCountry[PersonasDB[i].CountryCode]=countryInfo
// 		}

// 		countryCodePersona:=PersonasDB[i].CountryCode

// 		countryInfoPersona,ok:=mapaInfoCountry[countryCodePersona]

// 		personasExtendidas=append(personasExtendidas, PersonaExt{
// 			Persona: PersonasDB[i],
// 			CountryInfo: countryInfoPersona,
// 		})
// 	}

// 	jsonResponse, err := json.Marshal(personasExtendidas)
// 	if err != nil {
// 		http.Error(w, "Error al serializar personas", http.StatusInternalServerError)
// 		return
// 	}

// 	_, err = w.Write(jsonResponse)
// 	if err != nil {
// 		http.Error(w, "Error al escribir la respuesta", http.StatusInternalServerError)
// 		return
// 	}
// }
