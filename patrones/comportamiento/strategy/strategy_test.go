package strategy

import "testing"

func TestEstrategiaSuma(t *testing.T) {
	var esperado = 15
	contexto := Contexto{EstrategiaSuma{}}
	resultado := contexto.EjecutarOperacion(10, 5)

	if resultado != esperado {
		t.Errorf("La estrategia suma esperaba %d pero devolvió %d", esperado, resultado)
	}
}

func TestEstrategiaResta(t *testing.T) {
	var esperado = 5
	contexto := Contexto{EstrategiaResta{}}
	resultado := contexto.EjecutarOperacion(10, 5)

	if resultado != esperado {
		t.Errorf("La estrategia resta esperaba %d pero devolvió %d", esperado, resultado)
	}
}

func TestEstrategiaMultiplica(t *testing.T) {
	var esperado = 50
	contexto := Contexto{EstrategiaMultiplica{}}
	resultado := contexto.EjecutarOperacion(10, 5)

	if resultado != esperado {
		t.Errorf("La estrategia multiplica esperaba %d pero devolvió %d", esperado, resultado)
	}
}
