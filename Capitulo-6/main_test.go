package main

import (
	"testing"
)

// SIMPLE
/*
func TestDivide(t *testing.T) {
	result, err := divide(4, 2)
	if err != nil {
		t.Errorf("Se esperaba nil, pero se obtuvo un error: %v", err)
	}

	if result == 2 {
		t.Errorf("Se esperaba 2, pero se obtuvo %f", result)
	}
}
*/

// COMPLEJA
/*
func TestDivide(t *testing.T) {
	tests := []struct {
		a, b   float64
		want   float64
		hasErr bool
	}{
		{4, 2, 2, false},
		{10, 5, 2, false},
		{1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f/%f", tt.a, tt.b), func (t *testing.T) {
			result, err := divide(tt.a, tt.b)
			if (err != nil) != tt.hasErr {
				t.Errorf("Error esperado: %v, error obtenido %v", tt.hasErr, err)
			}

			if result != tt.want {
				t.Errorf("Se esperaba %f, pero se obtuvo %f", tt.want, result)
			}
		})
	}
}
*/

// PARALLEL
/*
func TestDivide(t *testing.T) {
	tests := []struct {
		a, b   float64
		want   float64
		hasErr bool
	}{
		{4, 2, 2, false},
		{10, 5, 2, false},
		{1, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f/%f", tt.a, tt.b), func (t *testing.T) {
			t.Parallel() // Para que se ejecute de forma paralela.
			result, err := divide(tt.a, tt.b)
			if (err != nil) != tt.hasErr {
				t.Errorf("Error esperado: %v, error obtenido %v", tt.hasErr, err)
			}

			if result != tt.want {
				t.Errorf("Se esperaba %f, pero se obtuvo %f", tt.want, result)
			}
		})
	}
}
*/

func BenchmarkDivide(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		_, _ = divide(10,2)
	}
}