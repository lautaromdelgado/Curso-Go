/*
	VIDEO: https://www.youtube.com/watch?v=qkYWCNEPPCo&ab_channel=RobertoMorais
	Minuto: 
*/

package main

import(
	"fmt"
	_ "net/http/pprof"
	"net/http"

	"time"
)

// Función para manejar las peticiones HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	// Simula una carga de trabajo
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "¡Hola, esta es una respuesta optimizada!")
}

func main() {
	// Configurar el handler para las peticiones en la ruta "/"
	http.HandleFunc("/", handler)

	// Inicia el servidor en el puerto 8080
	fmt.Println("El servidor se esta ejecutando en:")
	fmt.Println("http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}