package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(15, 12)
	esperado := 30

	if resultado != esperado {
		t.Errorf("Resultado inválido! Esperado: %d, Recebido: %d", esperado, resultado)
	}
}
