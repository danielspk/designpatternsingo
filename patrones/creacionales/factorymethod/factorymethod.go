package factorymethod

import "fmt"

// Interface Producto
type Entrevistador interface {
	RealizarPreguntas()
}

// Producto Concreto
type EntrevistadorIT struct{}

func (e *EntrevistadorIT) RealizarPreguntas() {
	fmt.Println("¿Nombre 5 patrones de diseño?")
}

// Producto Concreto
type EntrevistadorFinanzas struct{}

func (e *EntrevistadorFinanzas) RealizarPreguntas() {
	fmt.Println("¿Cuál es la alicuota del IVA?")
}

// Creador Interface
type RecursosHumanosInterface interface {
	LlamarEntrevistador() Entrevistador
}

// Creador Abstracto
type RecursosHumanos struct{}

func (rh *RecursosHumanos) TomarEntrevista(self RecursosHumanosInterface) {
	entrevistador := self.LlamarEntrevistador()
	entrevistador.RealizarPreguntas()
}

// Creador Concreto
type RecusosHumanosIT struct {
	*RecursosHumanos
}

func (rhi *RecusosHumanosIT) LlamarEntrevistador() Entrevistador {
	return &EntrevistadorIT{}
}

// Creador Concreto
type RecusosHumanosFinanzas struct {
	*RecursosHumanos
}

func (rhf *RecusosHumanosFinanzas) LlamarEntrevistador() Entrevistador {
	return &EntrevistadorFinanzas{}
}

// Test de uso del patrón
func TestPattern() {
	fmt.Println("El entrevisatador de IT pregunta:")
	recursosHumanosIT := &RecusosHumanosIT{&RecursosHumanos{}}
	recursosHumanosIT.TomarEntrevista(recursosHumanosIT)

	fmt.Println("\nEl entrevisatador de Finanzas pregunta:")
	recursosHumanosFinanzas := &RecusosHumanosFinanzas{&RecursosHumanos{}}
	recursosHumanosFinanzas.TomarEntrevista(recursosHumanosFinanzas)
}
