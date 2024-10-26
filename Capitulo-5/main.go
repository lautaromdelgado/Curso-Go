/*
	video: https://www.youtube.com/watch?v=tEBufGSmzAY&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=5&ab_channel=RobertoMorais
	minuto: 5:45
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
/*
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
*/
/* EJEMPLO SENCILLO */
// package main

// import "fmt"

// func main() {
// 	messages := make(chan string, 2) // Crear un canal con dos buffers, es decir 2 slots para guardar STRINGS
// 	messages <- "Hola"
// 	messages <- "Golang"
// 	fmt.Println(<-messages) // Output: Hola
// 	fmt.Println(<-messages) // Output: Golang
// }

/* EJEMPLO MÁS AVANZADO */

package main

import (
	"time"
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs{
		fmt.Printf("Trabajador %d empezó trabajo %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Trabajador %d terminó trabajo %d\n", id, j)
		results <- j * 2
	}
}

func main() {

}
