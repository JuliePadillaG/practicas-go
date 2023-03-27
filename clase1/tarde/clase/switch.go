package clase

import "fmt"

func main() {

	//Switch no requiere la palabra break, sale solo

	//Swith básico
	day := 1
	switch day {
	case 0:
		fmt.Println("Lunes")
	case 1:
		fmt.Println("Martes")
	default:
		fmt.Println("Desconocido")
	}

	//Switch sin expresion
	var edad uint8 = 18
	switch {
	case edad >= 150:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad")
	default:
		fmt.Println("Eres menor de edad")
	}

	//Switch con múltiples casos
	day1 := "domingo"

	switch day1 {
	case "lunes", "martes", "miercoles", "jueves", "viernes":
		fmt.Printf("%s es un día de la semana\n", day1)
	default:
		fmt.Printf("%s es un día del fin de la semana\n", day1)
	}

	//Switch con declaración corta
	switch day := "domingo"; day {
	case "lunes", "martes", "miercoles", "jueves", "viernes":
		fmt.Printf("%s es un día de la semana\n", day)
	default:
		fmt.Printf("%s es un día del fin de la semana\n", day)
	}

	//Switch con fallthrough
}