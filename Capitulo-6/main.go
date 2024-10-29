/*
	VIDEO: https://www.youtube.com/watch?v=umGvHL6VssY&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=6&ab_channel=RobertoMorais
	MINUTO: 4:00
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

// package main

// import (
// 	// "errors"
// 	"fmt"
// )

// var ErrNumberCero = errors.New("No se puede dividir por cero.")

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, ErrNumberCero
// 	}

// 	return a / b, nil
// }

// func main() {
// 	resultado, err := divide(10, 0)

// 	if err != nil {
// 		if errors.Is(err, ErrNumberCero) {
// 			fmt.Println("Error personalizado:", err)
// 			return
// 		}
// 	}
// 	fmt.Println("Resultado de la división es:", resultado)
// }

// DEFINIR UN ERROR MÁS DESCRIPTIVO
/*
type ErrorDetails struct {
	Arg  float64
	Prob string
}

func (e *ErrorDetails) Error() string {
	return fmt.Sprintf("%f - %s", e.Arg, e.Prob)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &ErrorDetails{x, "No puede ser un número negativo."}
	}

	return x * x, nil
}

func main() {
	_, err := sqrt(-4)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
*/

// FUNCIONES PANIC - RECOVER

/*
package main

import (
	"fmt"
)

func causePanic() {
	panic("Algo salió muy mal")
}

func main() {
	fmt.Println("Antes del panic")
	causePanic()
	fmt.Println("Esto no se imprimirá por consola")
}
*/

// RECOVER (Recuperar una goroutine que entra en panico)
// Solo es util en funciones diferidas.

/*
package main

import (
	"fmt"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado el panic", r)
	}
}

func riskyFunc() {
	defer handlePanic()
	panic("Se paró el programa inesperadamente..")
}

func main() {
	fmt.Println("Antes de la función riesgosa")
	riskyFunc()
	fmt.Println("Después de la función riesgosa.")
}
*/

// --- TESTING ---

package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por 0.")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10,0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("El resultado es:", result)
}