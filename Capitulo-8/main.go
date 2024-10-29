/*
	VIDEO: https://www.youtube.com/watch?v=fCSsJafKsjA&list=PLTcOzxm2NcYBBAZC-Ya_xqZ_eZ8i0o9NC&index=9&ab_channel=RobertoMorais
*/

package main

import (
	"encoding/json"
	"fmt"
	// "time"

	"net/http"
)

/*
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Accediste a la página principal")
}

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Esta es la página acerca de..")
// }

// MIDDLEWARE
func logginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Solicitud %s %s %s\n", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}

// MIDDLEWARE NUEVO
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer token123" {
			http.Error(w, "No Autorizado", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// METODO POST
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		fmt.Fprintf(w, "Nombre: %s\n", name)
		fmt.Fprintf(w, "Email: %s\n", email)
	} else {
		fmt.Fprintf(w, "Por favor, envía un formulario compatible.")
	}
}
*/

// MANEJAR FORMATO JSON

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:     "Lautaro",
		Lastname: "Delgado",
		Email:    "lautaromdelgado@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// http.HandleFunc("/home", homeHandler)
	// http.HandleFunc("/about", aboutHandler)
	// http.Handle("/form/data", logginMiddleware(http.HandlerFunc(formHandler)))
	// http.Handle("/", authMiddleware(logginMiddleware(http.HandlerFunc(homeHandler))))

	// JSON
	http.HandleFunc("/user", userHandler)

	fmt.Println("El Servidor se esta Ejecutando en el puerto :7001")
	http.ListenAndServe(":7001", nil)
}
