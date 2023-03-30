package main

import "fmt"

type Perro struct {
	Nombre string
	Edad int
	Raza string
	Peso float64
	Color []string
	marca string
	Owner Dueño
}

type Dueño struct {
	Documento int
	Nombre string
	NumeroContacto int
}

func main1() {

	//para inicializar dog pero todo en blanco
	//var dog Perro

	//para inicializar dog con campos definidos
	dog := Perro {
		Nombre: "Firulais",
		Edad: 2,
		Raza: "Criollo",
		Peso: 10.0,
		Color: []string{"Negro", "Blanco"},
		Owner: Dueño {
			Documento: 123,
			Nombre: "Julie",
			NumeroContacto: 123,
		},
	}

	

	fmt.Printf("Dog: %+v", dog)

	//Acceder a atributos del objeto
	fmt.Printf("\nDog: %+v", dog.Nombre)

	dog.marca = "Collie"
	fmt.Printf("\nDog: %+v\n", dog.marca)


}