package abstractfactory

import "fmt"

// Producto Abstracto Interface
type Puerta interface {
	VerMaterial() string
}

// Producto Concreto
type PuertaMadera struct{}

func (pm *PuertaMadera) VerMaterial() string {
	return "Madera"
}

// Producto Concreto
type PuertaMetal struct{}

func (pm *PuertaMetal) VerMaterial() string {
	return "Metal"
}

// Fábrica Abstracta Interface
type FabricaPuerta interface {
	ConstruirPuerta() Puerta
}

// Fábrica Concreta
type FabricaPuertaMadera struct{}

func (fpm *FabricaPuertaMadera) ConstruirPuerta() Puerta {
	return &PuertaMadera{}
}

// Fábrica Concreta
type FabricaPuertaMetal struct{}

func (fpm *FabricaPuertaMetal) ConstruirPuerta() Puerta {
	return &PuertaMetal{}
}

// Test de uso del patrón
func TestPattern() {
	fabricaPuertaMadera := &FabricaPuertaMadera{}
	puertaMadera := fabricaPuertaMadera.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMadera.VerMaterial())

	fabricaPuertaMetal := &FabricaPuertaMetal{}
	puertaMetal := fabricaPuertaMetal.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMetal.VerMaterial())
}
