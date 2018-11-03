package visitor

import "fmt"

// Visitante
type Visitante interface {
	VisitarSuperpoder(*ElementoSuperpoder)
	VisitarArma(*ElementoArma)
}

// Visitante Concreto
type VisitanteNivel1 struct{}

func (v *VisitanteNivel1) VisitarSuperpoder(elementoSuperpoder *ElementoSuperpoder) {
	elementoSuperpoder.Poder = "Rayo superpoderoso"
}

func (v *VisitanteNivel1) VisitarArma(elementoArma *ElementoArma) {
	elementoArma.Arma = "Espada de dos manos"
}

// Visitante Concreto
type VisitanteNivel0 struct{}

func (v *VisitanteNivel0) VisitarSuperpoder(elementoSuperpoder *ElementoSuperpoder) {
	elementoSuperpoder.Poder = "Rayo simple"
}

func (v *VisitanteNivel0) VisitarArma(elementoArma *ElementoArma) {
	elementoArma.Arma = "Espada de una mano"
}

// Elemento
type Elemento interface {
	Aceptar(Visitante)
}

// Elemento Concreto
type ElementoSuperpoder struct {
	Poder string
}

func (e *ElementoSuperpoder) Aceptar(visitante Visitante) {
	visitante.VisitarSuperpoder(e)
}

// Elemento Concreto
type ElementoArma struct {
	Arma string
}

func (e *ElementoArma) Aceptar(visitante Visitante) {
	visitante.VisitarArma(e)
}

// Test de uso del patrón
func TestPattern() {
	// desde visitar
	elementoArma0 := &ElementoArma{}
	elementoSuperpoder0 := &ElementoSuperpoder{}

	visitanteNivel0 := &VisitanteNivel0{}
	visitanteNivel0.VisitarArma(elementoArma0)
	visitanteNivel0.VisitarSuperpoder(elementoSuperpoder0)

	fmt.Printf("El visitante Nivel 0 tiene la siguiente arma de batalla: %s\n", elementoArma0.Arma)
	fmt.Printf("El visitante Nivel 0 tiene el siguiente superpoder: %s\n", elementoSuperpoder0.Poder)

	elementoArma1 := &ElementoArma{}
	elementoSuperpoder1 := &ElementoSuperpoder{}

	visitanteNivel1 := &VisitanteNivel1{}
	visitanteNivel1.VisitarArma(elementoArma1)
	visitanteNivel1.VisitarSuperpoder(elementoSuperpoder1)

	fmt.Printf("El visitante Nivel 1 tiene la siguiente arma de batalla: %s\n", elementoArma1.Arma)
	fmt.Printf("El visitante Nivel 1 tiene el siguiente superpoder: %s\n", elementoSuperpoder1.Poder)

	// desde aceptar
	visitanteNivel0 = &VisitanteNivel0{}
	elementoArma0 = &ElementoArma{}
	elementoArma0.Aceptar(visitanteNivel0)

	fmt.Printf("El elemento Arma aceptada por un visitante Nivel 0 es: %s\n", elementoArma0.Arma)
}
