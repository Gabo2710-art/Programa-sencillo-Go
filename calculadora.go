package main

import "fmt"

func main() {
	var num1, num2 float64

	fmt.Println("Ingrese el primer número:")
	fmt.Scanln(&num1)

	fmt.Println("Ingrese el segundo número:")
	fmt.Scanln(&num2)

	// Suma
	suma := num1 + num2
	fmt.Printf("La suma es: %.2f\n", suma)

	// Resta
	resta := num1 - num2
	fmt.Printf("La resta es: %.2f\n", resta)

	// Multiplicación
	multiplicacion := num1 * num2
	fmt.Printf("La multiplicación es: %.2f\n", multiplicacion)
}
