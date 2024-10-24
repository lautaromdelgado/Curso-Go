/*
	video: https://www.youtube.com/watch?v=tEBufGSmzAY&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=5&ab_channel=RobertoMorais
	minuto: 4:27
*/

// package main

// import (
// 	"fmt"
// 	// "time"
// 	"sync"
// )

// CONCURRENCIA

// GOroutines : Es un hilo de ejecución.

// func printMessage(timeout *sync.WaitGroup, message string) {
// 	defer timeout.Done() // Indica que la goroutine ha terminado
// 	fmt.Println(message)
// }

// func main() {
// 	var timeout sync.WaitGroup // Crear una instancia con el paque te sync.
// 	timeout.Add(1) // Añade una goroutine al WaitGroup
// 	go printMessage(&timeout, "¡Hola desde una goroutine!")
// 	timeout.Wait() // Espera a que todas las goroutines terminen

// 	// time.Sleep(time.Second) // Espera un segundo para que la GOroutine termine

// 	fmt.Println("¡Hola desde la función principal!")
// }

// --------------------------------------------------------------------------------------

// CHANNELS o CANALES

package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "Hola desde el canal :)"
	}()

	msg := <-messages
	fmt.Println(msg) // Muestro el canal "messages"
	// que se instancia en la linea 50
}
