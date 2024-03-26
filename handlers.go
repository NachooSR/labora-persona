package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//Enrutamientos

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
	id := insertarPersonaEnLaDb(aux)

	aux.ID = id

	encoder := json.NewEncoder(w)
	encoder.Encode(aux)

}

func getPersonas(w http.ResponseWriter, _ *http.Request) {
	// Aquí implementa la lógica para para obtener todas las personas
	var personasExtendidas []PersonaExt

	for i := 0; i < len(PersonasDB); i++ {
		countryInfo, err := getInfoCountry(PersonasDB[i].CountryCode)
		if err != nil {
			http.Error(w, "Error al obtener información del país", http.StatusInternalServerError)
			return
		}

		personasExtendidas = append(personasExtendidas, PersonaExt{
			Persona:     PersonasDB[i],
			CountryInfo: countryInfo,
		})
	}

	jsonResponse, err := json.Marshal(personasExtendidas)
	if err != nil {
		http.Error(w, "Error al serializar personas", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error al escribir la respuesta", http.StatusInternalServerError)
		return
	}
}

func getPerson(w http.ResponseWriter, r *http.Request) {

	idPersona := r.PathValue("id")

	idComoInt, _ := strconv.Atoi(idPersona)

	persona := ObtenerPersona(idComoInt)

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

	editarPersonaEnLaDb(personaAux.ID, personaAux.Edad, personaAux.Nombre, personaAux.Apellido, personaAux.CountryCode)

	encoder := json.NewEncoder(w)
	encoder.Encode(personaAux)

}

func deletePersona(w http.ResponseWriter, r *http.Request) {
	// Aquí implementa la lógica para eliminar una persona
}
