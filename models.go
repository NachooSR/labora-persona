package main

type Persona struct {
	ID int `json: "id"`
	Nombre string `json: "nombre"` 
	Apellido string `json: "apellido"`
	Edad int `json: "edad"`
	CountryCode string  `json: "countryCode"`
}
  
  // Comportamiento de la Persona..
func (p Persona) Validate() bool{
  
   if p.Apellido=="" || p.CountryCode=="" || p.Edad==0 || p.Nombre=="" || p.ID==0 {
	return false
   }
   return true
} 