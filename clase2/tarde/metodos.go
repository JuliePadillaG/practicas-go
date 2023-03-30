package main

import "fmt"

type Perro3 struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color []string `json:"color"`
	Owner Dueño3 `json:"owner"`
}

//La p referencia a la estructura que la contiene
func (p Perro3) Ladrar() {
	fmt.Println("Guau!")
}

func (p Perro3) Saludar() {
	fmt.Println("Guau!, Soy", p.Nombre)
}

//Si uso la función sin "*" estoy trabajando sobre una copia y el nombre en memoria no se cambia
//Pero si uso *, se cambia en memoria el valor
func (p *Perro3) Renombrar(new string) {
	p.Nombre = new
	fmt.Println("Guau!, Ahora Soy", p.Nombre)
}

type Dueño3 struct {
	Documento int
	Nombre string
	NumeroContacto int
}

func main4() {

	dog := Perro3 {
		Nombre: "Firulais",
		Edad: 2,
		Raza: "Criollo",
		Peso: 10.0,
		Color: []string{"Negro", "Blanco"},
		Owner: Dueño3 {
			Documento: 123,
			Nombre: "Julie",
			NumeroContacto: 123,
		},
	}

	dog.Ladrar()
	dog.Saludar()
	fmt.Println("Nombre antes", dog.Nombre)
	dog.Renombrar("Whisky")
	fmt.Println("Nombre después", dog.Nombre)
}