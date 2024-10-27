/*
	VIDEO: https://www.youtube.com/watch?v=umGvHL6VssY&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=6&ab_channel=RobertoMorais
	MINUTO: 1:41
*/

// TESTING

// package main

// import (
// 	"errors" // Paquete para manejar errores
// 	"fmt"    // Paqquete para mostrar por consola
// )

/*
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por cero.")
	}

	return a / b, nil
}

func main() {
	result, err := divide(4, 3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El valor de la división es:", result)
	}
}
*/

// ERRORES PERSONALIZADOS

package main

import(
	"fmt"
	"errors"
)

var ErrNumberCero = errors.New("No se puede dividir por cero.")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrNumberCero
	}

	return	a/b, nil
}

func main() {
	resultado, err := divide(10, 0)

	if err != nil {
		if errors.Is(err, ErrNumberCero) {
			fmt.Println("Error personalizado:", err)
			return
		}
	}
	fmt.Println("Resultado de la división es:", resultado)
}