package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 159)

	if total != 30 {
		t.Errorf("Resultado da soma está incorreto, Resultado: %d. Esperado: %d", total, 30)
	}
}
