package main

import (
	"database/sql"
	"os"

	//"encoding/json"
	"fmt"
	//"net/http"

	_ "github.com/lib/pq"
)

///El service es el intermediario entre la bdd y los handlers

type PersonaExt struct {
	Persona
	CountryInfo
}

var DB *sql.DB

func initDB() {
	var err error
	// Leer variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Usar las variables de entorno en la cadena de conexión
	connectionString := fmt.Sprintf("user=%s dbname=%s host=%s sslmode=disable password=%s port=%s", dbUser, dbName, dbHost, dbPassword, dbPort)
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error en la conexión a la base de datos")
		panic(err)
	}
}

// Inserta una persona ya cargada en la db, con un Id Autoincremental
func insertarPersonaEnLaDb(persona Persona) (int, error) {

	var idIngresado int

	errorcito := DB.QueryRow("INSERT into personas (nombre,apellido,edad,country_code) VALUES ($1, $2, $3, $4) RETURNING id;", persona.Nombre, persona.Apellido, persona.Edad, persona.CountryCode).Scan(&idIngresado)

	if errorcito != nil {
		return 0, errorcito
	}
	return idIngresado, nil
}

// Obtener una sola persona de la bdd buscada por su ID
func ObtenerPersona(id int) (PersonaExt, error) {

	///Persona auxiliar
	var personita_aux Persona

	err := DB.QueryRow("SELECT id,nombre,apellido,edad,country_code FROM personas WHERE id =$1", id).Scan(&personita_aux.ID, &personita_aux.Nombre, &personita_aux.Apellido, &personita_aux.Edad, &personita_aux.CountryCode)

	if err != nil {
		return PersonaExt{}, err
	}

	infoCountry, errorcito := getInfoCountry(personita_aux.CountryCode)
	if errorcito != nil {
		fmt.Println(errorcito)
		return PersonaExt{}, errorcito
	}

	return PersonaExt{
		Persona:     personita_aux,
		CountryInfo: infoCountry,
	}, nil
}

func editarPersonaEnLaDb(id, edad int, nombre, apellido, countryCode string) (*Persona, error) {

	var personaAux Persona

	///Reciclamos la funcion que tenemos
	personaAuxExtendida, err := ObtenerPersona(id)

	if err != nil {
		return nil, err
	}
	///Encontramos a la persona y la guardamos en una auxiliar para mantener sus datos
	personaAux = personaAuxExtendida.Persona

	//Si alguno de los parametros vino con valores lo reemplazamos en la variable auxiliar
	if nombre != "" {
		personaAux.Nombre = nombre
	}
	if apellido != "" {
		personaAux.Apellido = apellido
	}
	if edad != 0 {
		personaAux.Edad = edad
	}
	if countryCode != "" {
		personaAux.CountryCode = countryCode
	}

	_, err = DB.Exec("UPDATE personas SET nombre=$1 ,apellido=$2,edad=$3,country_code=$4 WHERE id=$5", personaAux.Nombre, personaAux.Apellido, personaAux.Edad, personaAux.CountryCode, id)
	if err != nil {
		return nil, err
	}
	return &personaAux, nil
}

func eliminarPersona(id int) error {

	_, err := DB.Exec("DELETE FROM personas WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
