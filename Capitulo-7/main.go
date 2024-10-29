package main

import (
	Calculadora "Capitulo-7/Go/mathutils"
	"fmt"
)

func main() {
	result := Calculadora.Add(10, 5)        // Llama a la función del archivo calculator.go
	result2 := Calculadora.Substract(10, 5) // Llama a la función del archivo calculator.go

	fmt.Printf("El resultado de Result es %d", result) // Muestra un mensaje por consola con el resultado
	fmt.Println()
	fmt.Printf("El resultado de Result2 es %d", result2) // Muestra un mensaje por consola con el resultado
}
