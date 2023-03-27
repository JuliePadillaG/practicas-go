package main

import "fmt"

/*Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que su sueldo es mejor a $100.000.

Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

//Solución presentada por los tutores


func main() {
	// Reglas:
	var (
		vEdad       int     = 22
		vEmpleado   bool    = true
		vAntiguedad int     = 1
		vSueldo     float64 = 100000.00
	)

	// Datos empleado:
	var (
		edad       = 23
		empleo     = true
		antiguedad = 1
		sueldo     = 200000.00
	)

	// Una forma seria realizar las validaciones mediante if y else if.
	// Pero otro forma quizas más elegante y corta seria.
	switch {
	case edad <= vEdad:
		fmt.Println("Debes ser mayor de 22 años de edad.")
	case empleo != vEmpleado:
		fmt.Println("Debes estar empleado.")
	case antiguedad < vAntiguedad:
		fmt.Println("Debes tener una antiguedad mayor a 1 año")
	case sueldo < vSueldo:
		fmt.Println("Otorgar credito sin interés ✅")
	default:
		fmt.Println("Otorgar credito con interés ✔️")
	}

}

//Solución hecha por mí

// func main() {

// 	var (
// 		edad int64 = 22
// 		empleado bool = true
// 		antiguedad int64 = 2
// 		sueldo float64 = 150000.0
// 	)

// 	switch {
// 	case edad >= 22:
// 		fmt.Println("Edad apta")
// 		fallthrough
// 	case empleado == true:
// 		fmt.Println("Apto")
// 		fallthrough
// 	case antiguedad >= 1:
// 		fmt.Println("apto")
// 	case sueldo > 100000:
// 		fmt.Println("No se cobrará interés")
// 	default:
// 		fmt.Println("No eres apto")
// 	}
// }