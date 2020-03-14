package chainofresponsibility

import "testing"

func TestMensajeAltaPrioridad(t *testing.T) {
	var esperado = "Procesando mensaje de alta prioridad: ALTA"
	manejadores := ReceptorBajaPrioridad{
		siguiente: ReceptorAltaPrioridad{},
	}
	resultado := manejadores.ProcesarMensaje(10, "ALTA")

	if resultado != esperado {
		t.Errorf("Se esperaba la siguiente respuesta del manejador '%s' pero se devolvió '%s'", esperado, resultado)
	}
}

func TestMensajeBajaPrioridad(t *testing.T) {
	var esperado = "Procesando mensaje de baja prioridad: BAJA"
	manejadores := ReceptorBajaPrioridad{
		siguiente: ReceptorAltaPrioridad{},
	}
	resultado := manejadores.ProcesarMensaje(3, "BAJA")

	if resultado != esperado {
		t.Errorf("Se esperaba la siguiente respuesta del manejador '%s' pero se devolvió '%s'", esperado, resultado)
	}
}
