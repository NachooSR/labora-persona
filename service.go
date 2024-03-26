package main



///El service es el intermediario entre la bdd y los handlers

var PersonasDB []Persona

//(C)Insertar Persona creada
//(R)Leer y devolver persona/personas (ext)
//(U)Actualizar
//(D)Eliminar


type PersonaExt struct{
	Persona
	CountryInfo
}



//Inserta una persona ya cargada en la db, con un Id Autoincremental
func insertarPersonaEnLaDb(persona Persona) int {
	persona.ID = len(PersonasDB) + 1
	PersonasDB = append(PersonasDB, persona)
	return persona.ID
}




// //Obtener personas
// func ObtenerPersonasDeLaDb() []PersonaExtendida {//DATA TRANSFER OBJECT : DTO
// 	llamadaAlServicioExterno()
// }


func ObtenerPersona(id int)Persona{
   for i := 0; i < len(PersonasDB); i++ {
	if PersonasDB[i].ID == id {
		return PersonasDB[i]
	}
   }

   return Persona{}
}


func editarPersonaEnLaDb(id, edad int, nombre, apellido, countryCode string)  {

	for i := 0; i < len(PersonasDB); i++ {
		if PersonasDB[i].ID == id {
			if nombre != "" {
				PersonasDB[i].Nombre = nombre
			}
			if apellido != "" {
				PersonasDB[i].Apellido = apellido
			}
			if edad != 0 {
				PersonasDB[i].Edad = edad
			}
			if countryCode != "" {
				PersonasDB[i].CountryCode = countryCode
			}
		}
	}
}



// func eliminarPersona(id int){

// }


