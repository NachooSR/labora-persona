package main

import (
	"fmt"
	"testing"
)


func cargarPersona2() Persona {
	var persona Persona = Persona{
		Nombre:      "Pedro",
		Apellido:    "Perez",
		Edad:        25,
		CountryCode: "Ar",
		ID: 10,
	}
	return persona
}


func TestValidate(t *testing.T) {
	
  counter:=0

  prueba_1:=cargarPersona2()
  
  prueba_2:=Persona{
	Nombre: "Carlos",
    Apellido: "Ramirez",
	Edad: 0,
	CountryCode: "Ar",
  }

  prueba_3:=Persona{
	Nombre: "Lionel",
  }

  ///Prueba 1
  resultado:=prueba_1.Validate()
  if(resultado==false){
   t.Fatalf("ERROR inesperado la persona contiene todos los datos")
   counter++
  }else{
	fmt.Println("Test 1- Pass")
  }
  
  ///Prueba 2
  resultado=prueba_2.Validate()
  if resultado {
	t.Fatalf("ERROR la persona posee datos invalidos")
  counter++
  }else{
	fmt.Println("Test 2- Pass")
  }

  ///Prueba 3
  resultado=prueba_3.Validate()
  if resultado {
	t.Fatalf("ERROR a la persona le faltan datos")
    counter++
  }else{
	fmt.Println("Test 3- Pass")
  }

  //**********************Pruebas finalizadas***************************//
  if counter==0{
	fmt.Println("Testeo exitoso, 0 errores encontrados")
  }

}