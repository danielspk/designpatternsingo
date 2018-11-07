package builder

import "fmt"

// Interface Constructor
type Constructor interface {
	EstablecerTamanio()
	Construir() *Hamburguesa
}

// Constructor Concreto
type ConstructorHamburguesaSimple struct {
	tamanio string
}

func (chs *ConstructorHamburguesaSimple) EstablecerTamanio() {
	chs.tamanio = "Simple"
}

func (chs *ConstructorHamburguesaSimple) Construir() *Hamburguesa {
	return &Hamburguesa{chs.tamanio}
}

// Constructor Concreto
type ConstructorHamburguesaDoble struct {
	tamanio string
}

func (chd *ConstructorHamburguesaDoble) EstablecerTamanio() {
	chd.tamanio = "Doble"
}

func (chd *ConstructorHamburguesaDoble) Construir() *Hamburguesa {
	return &Hamburguesa{chd.tamanio}
}

// Producto
type Hamburguesa struct {
	Tamanio string
}

// Director
type LocalComida struct{}

func (lc *LocalComida) ConstruirHamburguesa(constructor Constructor) *Hamburguesa {
	constructor.EstablecerTamanio()

	return constructor.Construir()
}

// Test de uso del patrón
func TestPattern() {
	localComida := &LocalComida{}

	hamburguesaA := localComida.ConstruirHamburguesa(&ConstructorHamburguesaSimple{})
	hamburguesaB := localComida.ConstruirHamburguesa(&ConstructorHamburguesaDoble{})

	fmt.Printf("Se solicito una hamburguesa: %s\n", hamburguesaA.Tamanio)
	fmt.Printf("Se solicito una hamburguesa: %s\n", hamburguesaB.Tamanio)
}
