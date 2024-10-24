package main

import "fmt"

// FUNCIONES
// func sayHello() { // Sin parametros
// 	fmt.Println("Hola desde Golang!")
// }

// func saludar(nombre string) string { // Con parametros y retorno
// 	return "Hola " + nombre + "!"
// }

// Primer MAIN
// func main() {
// 	// sayHello() // Llamando a la función
// 	// Output: Hola desde Golang!
// 	// fmt.Println(saludar("Lautaro Delgado"))

// 	// Función anonima (función que no tiene nombre)
// 	miNombre := "Lautaro Delgado"
// 	// func() {
// 	// 	fmt.Println("Desde una función anonima: " + miNombre)
// 	// }()

// 	func(nickname string) {
// 		fmt.Println("Llamando a " + nickname + " desde una función anonima con parametro")
// 	}(miNombre)
// }

// Segundo MAIN
// EJemplo un poco más completo de funciones anonimas
// func counter() func() int {
// 	var count int = 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func main() {
// 	var c = counter()

// 	fmt.Println(c()) // Primera impresión
// 	fmt.Println(c()) // Segunda impresión
// 	fmt.Println(c()) // Tercera impresión
// }

/* PUNTEROS */

// func cambiarValor(puntero *int) {
// 	*puntero = 42
// }

//  func main() {
// 	x := 10
// 	fmt.Println("Antes del cambio de valor: ",x)
// 	cambiarValor(&x)
// 	fmt.Println("Después del cambio de valor: ",x)
// }

// ESTRUCTURAS
type Address struct {
	City string
	Pais string
}

type Persona struct {
	Name string
	Surname string
	Age int
	Direccion Address
}

func main() {
	person := Persona{
		Name: "Lautaro",
		Surname: "Delgado",
		Age: 22,
		Direccion: Address{
			City: "Avellaneda",
			Pais: "Argentina",
		},
	}

	person2 := Persona{
		Name: "Sheila",
		Surname: "Biasoni",
		Age: 20,
		Direccion: Address{
			City: "Reconquista",
			Pais: "Argentina",
		},
	}

	person3 := Persona{
		Name: "Melania",
		Surname: "Nardelli",
		Age: 19,
		Direccion: Address{
			City: "Roma",
			Pais: "Italia",
		},
	}

	fmt.Println("En la estructura Persona, los datos cargados son:")
	fmt.Println("El usuario se llama ",person.Name, person.Surname, " y tiene ",person.Age," años de edad, viviendo en ",person.Direccion.City," de ",person.Direccion.Pais)
	fmt.Println("El usuario se llama ",person2.Name, person2.Surname, " y tiene ",person2.Age," años de edad, viviendo en ",person2.Direccion.City," de ",person2.Direccion.Pais)
	fmt.Println("El usuario se llama ",person3.Name, person3.Surname, " y tiene ",person3.Age," años de edad, viviendo en ",person3.Direccion.City," de ",person3.Direccion.Pais)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("La instacia Person almacena:", person)
	fmt.Println("La instacia Person2 almacena:", person2)
	fmt.Println("La instacia Person3 almacena:", person3)
}