package main

import (
	"fmt"
	"encoding/json"
)

type Perro1 struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color []string `json:"color"`
	Owner Dueño1 `json:"owner"`
}

type Dueño1 struct {
	Documento int
	Nombre string
	NumeroContacto int
}

func main2() {

	//para inicializar dog pero todo en blanco
	//var dog Perro

	//para inicializar dog con campos definidos
	dog := Perro1 {
		Nombre: "Firulais",
		Edad: 2,
		Raza: "Criollo",
		Peso: 10.0,
		Color: []string{"Negro", "Blanco"},
		Owner: Dueño1 {
			Documento: 123,
			Nombre: "Julie",
			NumeroContacto: 123,
		},
	}

	//Marshall -> Devuelve una representación en JSON de una estructura, en bytes 
	//Marshall usa un concepto de reflexión y necesita la dirección en memoria del objeto para entrar y leerla
	b, err := json.Marshal(&dog)
	if err != nil {
		panic(err)
	}

	//se castea a string para que no devuelva el escritos los bytes
	//si no queremos enviar un dato del struct al json que se va a consumir lo escribimos en minúscula
	fmt.Printf("\njson.Marshall: %+v", string(b))
}