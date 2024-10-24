// Video del capitulo 4: https://www.youtube.com/watch?v=qjZKTSYKUeM&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=4&ab_channel=RobertoMorais
// Minuto 07:46

// METODOS

package main

import "fmt"

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimetro() float64 {
// 	r.Width = 50
// 	return 2 * (r.Width + r.Height)
// }

// func (r *Rectangle) Scale(factor float64) {
// 	r.Width *= factor
// 	r.Height *= factor
// }

// func main() {
// 	// rectangulo := Rectangle{Width: 10, Height: 10}
// 	// fmt.Println("Area del rectangulo:", rectangulo.Area()) // Output: 100
// 	/* -------------------------------------------------------------------------- */
// 	// r := Rectangle{Width: 100, Height: 5}
// 	// rec := Rectangle{Width: 500, Height: 100}
// 	// fmt.Println("Perimetro del rectangulo:", r.Perimetro()) // Output: 30
// 	// fmt.Println("El perimetro del rectangulo de 'rec' es:", rec.Perimetro())
// 	/* -------------------------------------------------------------------------- */
// 	r := Rectangle{Width: 50, Height: 100}
// 	r.Scale(2.5)
// 	fmt.Println("Scale:", r)
// }

// Receivers:
/*
	De Valor: Cuando no se necesite cambiar el valor original.
	De Puntero: Cuando se necesite cambiar el valor original.
*/

// INTERFACES

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	printArea(5)
}