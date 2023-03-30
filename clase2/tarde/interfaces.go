package main

import "fmt"

type Perro4 struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Color []string `json:"color"`
	Owner Dueño4 `json:"owner"`
}

func (p Perro4) Ladrar() {
	fmt.Println("Guau!")
}

func (p Perro4) Saludar() {
	fmt.Println("Guau!, Soy", p.Nombre)
}

//Si uso la función sin "*" estoy trabajando sobre una copia y el nombre en memoria no se cambia
//Pero si uso *, se cambia en memoria el valor
func (p *Perro4) Renombrar(new string) {
	p.Nombre = new
	fmt.Println("Guau!, Ahora Soy", p.Nombre)
}

func (p Perro4) Dormir() {
	fmt.Println("Estoy durmiendo, soy un perro")
}

func (p Perro4) Comer() {
	fmt.Println("Estoy comiendo, soy un perro")
}

type Dueño4 struct {
	Documento int
	Nombre string
	NumeroContacto int
}

type Gato struct {
	Nombre string `json:"nombre"`
	Edad int `json:"edad"`
	Raza string `json:"raza"`
	Peso float64 `json:"peso"`
	Caracter string `json:"caracter"`
}

func (g Gato) Dormir() {
	fmt.Println("Estoy durmiendo, soy un gato")
}

func (g Gato) Comer() {
	fmt.Println("Estoy comiendo, soy un gato")
}

//Definición de la interfaz
//Dentro de la interfaz se definen los métodos que se tienen que cumplir (contrato)
type Animal interface {
	Dormir()
	Comer()
}

func NewAnimal(tipo string) Animal {
	if tipo == "Perro"{
		return &Perro4{}
	} else if tipo == "Gato" {
		return &Gato{}
	}
	return nil
}

func main() {
	a := NewAnimal("Gato")
	a.Comer()
	a.Dormir()
}