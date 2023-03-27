package main

// Ejercicio 1 - Imprimí tu nombre
// 1. Crea una aplicación donde tengas como variable tu nombre y dirección.
// 2. Imprime en consola el valor de cada una de las variables.

import "fmt"


var(
	name string = "Julie"
	direccion string = "Calle falsa 123"
)

func main() {
	fmt.Printf("Mi nombre es %s y vivo en %s\n", name, direccion)
}

