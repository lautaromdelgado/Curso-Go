package main

import(
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Crea una peticion GET simulada
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("No se pudo crear la petición %v", err)
	}

	// Graba la respuesta en el handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	// Llamar al handler con la peticion simulada
	handler.ServeHTTP(rr, req)

	// Verifica que el código de estado es 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código incorrecto: obtuvo %v", status)
	}

	// Verifica que la respuesta contiene el mensaje esperado
	expected := "¡Hola, esta es una respuesta optimizada!\n"

	if rr.Body.String() != expected {
		t.Errorf("Handler devolvio un cuerpo inesperado: obtuvo %v", rr.Body.String())
	}
}