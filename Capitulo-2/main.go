package main

import "fmt"

// DECLARACION DE VARIABLES
// func main() {
// 	// Primera manera de declarar una variable
// 	var nombre string = "Lautaro Delgado"
// 	// Segunda manera de declarar una variable
// 	var nombre2 = "Iván Escobar"
// 	// Tercera manera de declarar una variable
// 	nombre3 := "Lautaro Martínez"

// 	fmt.Println(nombre)  // Salida por terminal
// 	fmt.Println(nombre2) // Salida por terminal
// 	fmt.Println(nombre3) // Salida por terminal
// }

// Flujo de Control con IF - ELSE IF - ELSE
// func main() {
// 	var nombre string = "Diego"
// 	var edad int = -1

// 	if edad >= 18 {
// 		fmt.Println(nombre, "puedes pasar al boliche") // += 18
// 	} else if edad < 18 && edad > 0 {
// 		fmt.Println(nombre, "no puedes pasar al boliche") // - 18
// 	} else {
// 		fmt.Println("¡No se aceptan números negativos!") // Numero negativo
// 	}
// }

// Flujo de Control con SWITCH
// func main() {
// 	dia := "Domingo"

// 	switch dia {
// 	case "Lunes":
// 		fmt.Println("Inicio de semana laboral")
// 		break
// 	case "Martes":
// 		fmt.Println("Seguimos con el trabajo")
// 		break
// 	case "Miercoles":
// 		fmt.Println("Estamos a mitad de semana laboral")
// 		break
// 	case "Jueves":
// 		fmt.Println("Estamos cerca del fin de semana")
// 		break
// 	case "Viernes":
// 		fmt.Println("Inicio del fin de semana")
// 		break
// 	default:
// 		fmt.Println("Error al interpretar el dia")
// 		break
// 	}
// }

// Bucle FOR

// func main() {
// 	/* Bucle FOR normal*/
// 	// for i := 0; i < 11; i++ {
// 	// 	fmt.Println("Ahora valgo", i)
// 	// }

// 	/* Bucle FOR como WHILE*/
// 	var i int = 0
// 	for i <= 10 {
// 		fmt.Println("Soy I y valgo:", i)
// 		i++
// 	}
// }

// func main() {
// 	i := 0 // Inicialización de I

// 	for i <= 100 {
// 		fmt.Println("Valgo:", i)
// 		i++
// 		if i < 51 {
// 			continue
// 		}
// 		break
// 	}
// }

/* ESTRUCTURA DE DATOS */
// Arrays - Slices - Maps

// ARRAY
/*
func main() {
	var miArray [5]int // Inicialización del array
	contador := 0 // Inicialización del contador

	for contador < 5 { // Condición de parada
		miArray[contador] = contador * 3 // Asignación de valores
		contador++ // Contador de iteraciones
	}
	// Muestro los valores del array por terminal
	fmt.Println(miArray) // output : valores del array
}
*/

// SLICE
// func main() {
// 	var miSlice = []string{}                       // Inicialización del SLICE
// 	miSlice = append(miSlice, "Lautaro", "Sheila") // Agregación de valores
// 	otroSlice := append(miSlice, "Melania")        // Agregación de valores

// 	fmt.Println(miSlice)   // Salida por terminal
// 	fmt.Println(otroSlice) // Salida por terminal
// 	fmt.Println(miSlice)   // Salida por terminal
// }

// MAPS (MAPAS) - Estilo objetos en Javascript (Clave : Valor)
func main() {
	edades := make(map[string]int)

	edades["Lautaro"] = 22
	edades["Sheila"] = 20
	edades["Melania"] = 19

	for nombre, edad := range edades {
		fmt.Println(nombre, "tiene", edad, "años.")
	}
}