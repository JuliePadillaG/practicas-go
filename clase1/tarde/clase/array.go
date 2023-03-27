package clase

import (
	"fmt"
	)

func Ejecutarclase() {
	//var myArray [2]int
	//Declarar un array con valores predefinidos
	// Los tres punticos significan "no tengo el tamaño definido, quiero que lo deduzcas de los valores iniciales"
	var myArray = [...]int{1,2,3}
	fmt.Println(myArray)

	//Para crear un Slice
	a := make([]int, 2)
	fmt.Println(a)
	slc := myArray[0:2]

	//Otra forma para definir un slice sin tamaño definido desde el inicio
	var slc2 []int
	fmt.Printf("Este es el silce sin tamaño definido %d\n", slc2)

	//Para obtener rango de valores
	fmt.Println(a[:])

	//Capacidad y tamaño de slices
	fmt.Println(cap(slc), len(slc))

	//Agregar valores a un slice
	slc = append(slc, 1, 2, 3)
	fmt.Println(slc)
	fmt.Println(cap(slc), len(slc))
}