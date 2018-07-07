package chainofresponsability

// Interface
type Receptor interface {
	ProcesarMensaje(int, string) string
}

// Receptor de Alta Prioridad
type ReceptorAltaPrioridad struct {
	siguiente Receptor
}

func (rap ReceptorAltaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
	if prioridad >= 5 {
		return "Procesando mensaje de alta prioridad: " + mensaje
	}

	if rap.siguiente != nil {
		return rap.siguiente.ProcesarMensaje(prioridad, mensaje)
	}

	return ""
}

// Receptor de Baja Prioridad
type ReceptorBajaPrioridad struct {
	siguiente Receptor
}

func (rbp ReceptorBajaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
	if prioridad < 5 {
		return "Procesando mensaje de baja prioridad: " + mensaje
	}

	if rbp.siguiente != nil {
		return rbp.siguiente.ProcesarMensaje(prioridad, mensaje)
	}

	return ""
}
