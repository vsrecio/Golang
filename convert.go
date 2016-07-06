package main

import "fmt"

func main() {
	fmt.Print("Entra la temperatura en Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9

	fmt.Println(celcius)

}
