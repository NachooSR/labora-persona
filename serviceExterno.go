package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

///Aca traemos los datos de la Api externa



type CountryRTA []struct {
	Name         Name             `json:"name"`        
	Timezones    []string         `json:"timezones"` 
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Flags        Flags                  `json:"flags"`       
}


type Currencies struct {
	Ars Ars `json:"ARS"`
}

type Ars struct {
	Name   string `json:"name"`  
	Symbol string `json:"symbol"`
}


type Flags struct {
	PNG string `json:"png"`
}

type Name struct {
	Common string `json:"common"`    
}

type CountryInfo struct{
     Name string
	 Timezone string
	 Currency string
	 Flag string
}


//Funcion que recibe el codigo de la persona y devuelve la informacion del mismo Pais
func getInfoCountry(codeCountry string)(CountryInfo,error){
     
	url:= fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s",codeCountry)
	resp,err:=http.Get(url)

	if(err!=nil){
		return CountryInfo{}, fmt.Errorf("ERROR: %s",resp.Status)
	}

	var countryRTA CountryRTA

	///Decodeamos la info que nos trajo la Api en formato Json
	err=json.NewDecoder(resp.Body).Decode(&countryRTA)

	if err!=nil {
		return CountryInfo{},err
	}
    
	auxCountry:= countryRTA[0]

	CountryCorrespondiente:= CountryInfo{
         Name: auxCountry.Name.Common,
		 Timezone: auxCountry.Timezones[0],
		 Currency: auxCountry.Currencies["COP"].Name,
        Flag: auxCountry.Flags.PNG,
	}

	return CountryCorrespondiente,nil
}
