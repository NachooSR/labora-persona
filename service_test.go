package main

import (
	"fmt"
	"testing"
)

func cargarPersona() Persona {
	var persona Persona = Persona{
		Nombre:      "Pedro",
		Apellido:    "Perez",
		Edad:        25,
		CountryCode: "Ar",
	}
	return persona
}

func Test_initDB(t *testing.T) {

	initDb()

	err := DB.Ping()
	if err != nil {
		t.Fatalf("Error al conectar la Base de datos: %v", err)
		return
	}
	fmt.Println("Base de datos conectada exitosamente")
}

func Test_insertarPersonaEnLaDb(t *testing.T) {

	initDb()

	personaAux := cargarPersona()

	idIngresado, err := insertarPersonaEnLaDb(personaAux)

	//En mi db hay 3 registros ingresados por lo que la siguiente persona tiene que tener Id=4
	if err != nil || idIngresado != 4 {
		t.Fatalf("La persona no se pudo ingresar, ocurrio un error: %v", err)
		return
	}

	fmt.Println("La persona se ingreso correctamente tiene el id: ", idIngresado)

	eliminarPersona(idIngresado)

	//Esto es para reiniciar el id, sino se crean registros, se eliminan y los id que se fueron borrando van quedando
	///Por lo que queda algo como 1,2,3,8,9...
	_, err = DB.Exec("ALTER SEQUENCE personas_id_seq RESTART WITH 4")
	if err != nil {
		panic(err)
	}

}

func Test_editarPersonaEnLaDb(t *testing.T) {

	initDb()

	///Ingresaremos un registro auxiliar
	///Lo editamos, hacemos un select para ver que se ingreso y se modifico
	///Luego lo eliminamos

	personaAuxiliar := cargarPersona()

	id_Prueba, err := insertarPersonaEnLaDb(personaAuxiliar)

	if err != nil {
		t.Fatalf("ERROR: %v", err)
		return
	}

	///Persona ingresada: Pedro Perez 25 Ar
	///Persona modificada:Pedro Perez 26 Br
	_, err = editarPersonaEnLaDb(id_Prueba, 26, "", "", "Br")
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}

	///La seleccionamos de la db para ver si quedo editada
	err = DB.QueryRow("SELECT id,nombre,apellido,edad,country_code FROM personas WHERE id=$1", id_Prueba).Scan(&personaAuxiliar.ID, &personaAuxiliar.Nombre, &personaAuxiliar.Apellido, &personaAuxiliar.Edad, &personaAuxiliar.CountryCode)

	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}

	///Confirmamos los valores
	if personaAuxiliar.CountryCode != "Br" || personaAuxiliar.Edad != 26 {
		t.Fatalf("ERROR***** La persona no se edito correctamente")
		return
	}

	///La eliminamos y reseteamos el id
	eliminarPersona(id_Prueba)
	_, err = DB.Exec("ALTER SEQUENCE personas_id_seq RESTART WITH 4")
	if err != nil {
		panic(err)
	}
	fmt.Println("Persona editada correctamente")
}

func Test_eliminarPersona(t *testing.T) {
	initDb()

	persona := cargarPersona()
	id_Prueba, err := insertarPersonaEnLaDb(persona)

	if err != nil {
		t.Fatalf("ERROR: %v", err)
		return
	}

	err = eliminarPersona(id_Prueba)
	if err != nil {
		t.Fatalf("ERROR: %v", err)
		return
	}

	existe, errorcito := existeId(id_Prueba)
	if errorcito != nil {
		t.Fatalf("ERROR: %v", err)
		return
	}
	if existe {
		t.Fatalf("La persona no se elimino")
		return
	}

	fmt.Println("Persona eliminada correctamente")
}

func existeId(id int) (bool, error) {
	var existe bool

	initDb()

	err := DB.QueryRow("SELECT EXISTS (SELECT 1 FROM personas WHERE id=$1)", id).Scan(&existe)
	if err != nil {
		return false, err
	}

	return existe, nil
}
