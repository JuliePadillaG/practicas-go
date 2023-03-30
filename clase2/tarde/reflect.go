package main

import (
	"reflect"
	"fmt"
)

type Perro2 struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color []string `json:"color"`
	Owner Dueño2 `json:"owner"`
}

type Dueño2 struct {
	Documento int
	Nombre string
	NumeroContacto int
}

func main3() {

	dog := Perro2 {
		Nombre: "Firulais",
		Edad: 2,
		Raza: "Criollo",
		Peso: 10.0,
		Color: []string{"Negro", "Blanco"},
		Owner: Dueño2 {
			Documento: 123,
			Nombre: "Julie",
			NumeroContacto: 123,
		},
	}

	//reflect hace una copia en tiempo real de la estructura de datos por eso no necesita como argumento la dirección en memoria
	dogR := reflect.TypeOf(dog)

	for i := 0; i < dogR.NumField(); i++ {
		field := dogR.Field(i)
		//Dentro del get podemos poner las etiquetas personalizadas, en este caso no hay y se pone la que tiene json
		fmt.Printf("\nCampo: %s, tiene %s\n", field.Name, field.Tag.Get("json"))
	}

}