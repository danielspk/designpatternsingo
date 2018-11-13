package facade

import "fmt"

// Subsistema
type BuscadorHotel struct{}

func (bh *BuscadorHotel) BuscarHabitacion(entrada string, salida string) []string {
	return []string{
		"Hotel A con Habitación Classic - Entrada " + entrada + ", Salida " + salida + " - $500.00",
		"Hotel B con Habitación Deluxe - Entrada " + entrada + ", Salida " + salida + " - $750.00",
	}
}

// Subsistema
type BuscadorAvion struct{}

func (ba *BuscadorAvion) BuscarPasaje(salida string, regreso string) []string {
	return []string{
		"Aerolinea A - Clase Turista - Salida " + salida + ", Regreso " + regreso + " - $2400.00",
		"Aerolinea A - Clase Ejecutiva - Salida " + salida + ", Regreso " + regreso + " - $3200.00",
		"Aerolinea B - Clase Ejecutiva - Salida " + salida + ", Regreso " + regreso + " - $3800.00",
	}
}

// Fachada
type AgenciaViaje struct{}

func (av *AgenciaViaje) BuscarPaquete(desde string, hasta string) string {
	buscadorHoteles := &BuscadorHotel{}
	buscadorAvion := &BuscadorAvion{}

	resultadosHabitaciones := buscadorHoteles.BuscarHabitacion(desde, hasta)
	resultadosPasajes := buscadorAvion.BuscarPasaje(desde, hasta)

	respuesta := "Hoteles disponibles:\n"

	for _, habitacion := range resultadosHabitaciones {
		respuesta = respuesta + " - " + habitacion + "\n"
	}

	respuesta = respuesta + "\nAerolineas disponibles:\n"

	for _, pajase := range resultadosPasajes {
		respuesta = respuesta + " - " + pajase + "\n"
	}

	return respuesta
}

// Test de uso del patrón
func TestPattern() {
	agencia := &AgenciaViaje{}

	fmt.Printf("%s\n", agencia.BuscarPaquete("2018-12-01", "2018-12-08"))
}
