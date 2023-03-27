package clase

import "fmt"

//Definición de map
func clase() {
	//Inicializa con valores
	myMap1 := map[string]int{}
	//Inicializa vacío
	myMap2 := make(map[string]string)

	fmt.Println(myMap1)
	fmt.Println(myMap2)

	//Declarar un map y acceder a valores
	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}
	fmt.Println(students["Benjamin"])

	//Agregar nuevo elemento
	students["Brenda"] = 16
	fmt.Println(students)

	//Eliminar elemento
	delete(students, "Benjamin")
	fmt.Println(students)

	//Recorrer elementos de un map
	for key, element := range students {
		fmt.Println("Key: ", key, "=>", "Element: ", element)
	}

	//Para comprobar si existe o no esa clave
	x, ok := students["Pedro"]
	fmt.Println(x, ok)
	if !ok {	
		fmt.Printf("El valor %d no existe\n", x)
	}
}

