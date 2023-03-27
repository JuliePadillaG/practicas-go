package main

import "fmt"


// Ejercicio 2 - Clima

// Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares. 
// Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprime los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

//Solución propuesta

func main() {
	var temperatura, humedad, pression float32 = 16, 80, 1000.2

	fmt.Println("Reporte del clima: 🌥️")
	fmt.Printf("Temperatura: %f C°\n", temperatura)
	fmt.Printf("Humedad: %.0f %%\n", humedad)                // %% para escapar el porcentaje
	fmt.Printf("Presión Atmosferica: %.2f hPa°\n", pression) //o hectopascal (hPa).
}

//Solución alcanzada
// var (
// 	temperatura float64 = 10.0
// 	humedad float64 = 20
// 	presion float64 = 3
// )

// func main() {
// 	fmt.Printf("La temperatura es %.1f %%, la humedad %2f y la presión %2f\n", temperatura, humedad, presion)
// }
