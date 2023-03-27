package clase

import "fmt"

func clase() {
	//Standard for
	sum := 1
	for i := 0; i < 100; i++ {
		sum += i
	}

	//For range
	frutas := []string{"manzana", "banana", "pera"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	//Bucle infinito
	sum = 1
	for {
		sum++
	}

	//Bucle while
	sum = 2
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	//Romper un bucle
	sum = 0
	for {
		sum++
		if sum >= 1000 {
			break
		}
	}
	fmt.Println(sum)

	//Saltar a la siguiente iteración
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}
}