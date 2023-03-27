package main

import "fmt"

	/*Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa,
ayuda  a imprimir la edad de Benjamin.

Por otro lado también es necesario:
	1. Saber cuántos de sus empleados son mayores de 21 años.
	2. Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	3. Eliminar a Pedro del mapa.
*/

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	//Edad de Benjamín
	fmt.Printf("La edad de Benjamín es %d años\n", employees["Benjamin"])

	//Mayores de 21
	for key, element := range employees {
		if (element > 21) {
			fmt.Printf("%s Es mayor de 21 años \n", key)
		}
	}

	//Agregar a Federico 25 años
	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar a Pedro
	delete(employees, "Pedro")
	fmt.Println(employees)
}