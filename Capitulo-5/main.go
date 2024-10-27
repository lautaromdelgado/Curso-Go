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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("Trabajador %d empezó trabajo %d\n", id, j)
// 		time.Sleep(time.Second)
// 		fmt.Printf("Trabajador %d terminó trabajo %d\n", id, j)
// 		results <- j * 2
// 	}
// }

// func main() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for a := 1; a <= 5; a++ {
// 		fmt.Println("Resultado:", <-results)
// 	}
// }

// SELECT

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second) // Va a almacenar primero el valor en canal 1
// 		ch1 <- "Canal uno."
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second) // Va a almacenar último el valor en el canal 2
// 		ch2 <- "Canal dos."
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Ricibido:", msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println("Recibido:", msg2)
// 		}
// 	}
// }

/*
	SELECT para implementar TimeOut, esperando un tiempo
	especifico si no se recibe ningún valor.
*/

package main

// import(
// 	"fmt"
// 	"time"
// )

// func main() {
// 	channel := make(chan string)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		channel <- "Recibido en el canal por Goroutine"
// 	}()

// 	select {
// 	case res := <- channel:
// 		fmt.Println(res)
// 	case <-time.After(time.Second * 3):
// 		fmt.Println("Goroutine no recibida..")
// 	}
// }

// import (
// 	"fmt"
// 	"sync"
// )

// type Counter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// func (c *Counter) Value() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value
// }

// func main() {
// 	c := &Counter{}
// 	var wg sync.WaitGroup

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()
// 			c.Increment()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("Final del Counter:", c.Value())
// }

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	for j := 0; j < 1000; j++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println("Valor actual de la goroutine:", c.Value())
		}()
	}

	wg.Wait()

	fmt.Println("Final del Counter:", c.Value())
}
