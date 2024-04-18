package main

import (
	"database/sql"
	"fmt"
	"testing"
)

func Test_editarPersonaEnLaDb(t* testing.T){
	
	///Conectamos la db
	var err error
	DB,err=sql.Open("postgres","user=postgres dbname=personas password=bocajuniors,2022 host=localhost sslmode=disable port=5432")
	if err!=nil {
		t.Fatalf("Error al conectar Db: %v",err)
		return
	}

	fmt.Println("BDD conectada exitosamente")
	defer DB.Close()

	///Inicio transaccion que va ser cancelada con un rollback
	tx,errorcito:=DB.Begin()
	
	if errorcito!=nil {
         t.Fatalf("Error al iniciar transaccion: %v",errorcito)
		 return
	}

	fmt.Println("Transaccion iniciada exitosamente")

	defer tx.Rollback()///Que no se comiteen los cambios y asi no dejar huella en la bd 

	///Aca escribo la transaccion
	///1-Insert persona
	///2-Editarla
	///3-Verificar que se edito
	///Rollback
	
	var personita Persona=Persona{
		Nombre: "Juan",
		Apellido: "Perez",
		Edad: 19,
		CountryCode: "Ar",
	}
    var idPersona int 
	
	idPersona,err=insertarPersonaEnLaDb(personita)
    if err!=nil {
		t.Fatalf("Error al insertar la persona: %v",err)
		return
	}

	fmt.Println("Persona ingresada exitosamente con id: ",idPersona)


	//aux
	var personaAux *Persona
	personaAux,err= editarPersonaEnLaDb(idPersona,0,"Ignacio","Reyna","")

	if err!=nil{
		t.Fatalf("Error al editar %v",err)
		return
	}
	

	if personaAux.CountryCode!="Ar" || personaAux.Nombre!="Ignacio" || personaAux.Apellido !="Reyna" ||personaAux.Edad!=19 {
		t.Fatalf("La persona no se edito en la db")	
		return
	}

	fmt.Println("Persona editada correctamente")

   err=eliminarPersona(idPersona)
   if(err!=nil){
	t.Fatalf("Error al eliminar: %v",err)
	return
   }

   fmt.Println("Persona eliminada correctamente")
}
