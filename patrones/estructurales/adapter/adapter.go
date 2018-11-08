package adapter

import "fmt"

// Objetivo
type Gerrero interface {
	UsarArma() string
}

type Elfo struct{}

func (e *Elfo) UsarArma() string {
	return "atacando con arco y flecha"
}

// Adaptable
type GerreroMagico interface {
	UsarMagia() string
}

type Mago struct{}

func (m *Mago) UsarMagia() string {
	return "atacando con magia"
}

// Adaptador
type MagoAdaptador struct {
	gerrero GerreroMagico
}

func (ma *MagoAdaptador) UsarArma() string {
	return ma.gerrero.UsarMagia()
}

// Cliente
type Jugador struct {
	guerrero Gerrero
}

func (j *Jugador) Atacar() string {
	return j.guerrero.UsarArma()
}

// Test de uso del patrón
func TestPattern() {
	jugadorA := &Jugador{&Elfo{}}
	fmt.Printf("Jugador A: %s\n", jugadorA.Atacar())

	jugadorB := &Jugador{&MagoAdaptador{&Mago{}}}
	fmt.Printf("Jugador B: %s\n", jugadorB.Atacar())
}
