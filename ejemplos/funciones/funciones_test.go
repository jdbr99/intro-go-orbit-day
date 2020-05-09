package cuadrado

import (
	"testing"
)

func TestCuadrado(t *testing.T) {
	if val := Cuadrado(2); val != 4 {
		t.Errorf("Error! Se esperaba 4 y se obtuvo %d", val)
	}
}

func TestCuadradoTabla(t *testing.T) {
	var testsTable = []struct {
		entrada int
		salida  int
	}{
		{2, 4},
		{3, 9},
		{4, 16},
		{5, 25},
		{40, 1600},
	}

	for i, prueba := range testsTable {
		t.Logf("Corriendo prueba no.%d\n", i)
		if res := Cuadrado(prueba.entrada); res != prueba.salida {
			t.Errorf("Error! Esperaba %d, pero obtuve %d\n", prueba.salida, res)
		}
	}
}
