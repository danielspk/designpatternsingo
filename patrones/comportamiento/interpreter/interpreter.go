package interpreter

import "fmt"

// Contexto
type Contexto struct {
	Palabra string
}

// Interface
type ExpresionAbstracta interface {
	Interpretar(Contexto) bool
}

// Expresion Terminal
type ExpresionTerminal struct {
	Palabra string
}

func (et *ExpresionTerminal) Interpretar(contexto Contexto) bool {
	return contexto.Palabra == et.Palabra
}

// Expresion No Terminal
type ExpresionOR struct {
	expresionA ExpresionAbstracta
	expresionB ExpresionAbstracta
}

func (eo *ExpresionOR) Interpretar(contexto Contexto) bool {
	return eo.expresionA.Interpretar(contexto) || eo.expresionB.Interpretar(contexto)
}

// Expresion No Terminal
type ExpresionAND struct {
	expresionA ExpresionAbstracta
	expresionB ExpresionAbstracta
}

func (ea *ExpresionAND) Interpretar(contexto Contexto) bool {
	return ea.expresionA.Interpretar(contexto) && ea.expresionB.Interpretar(contexto)
}

// Test de uso del patrón
func TestPattern() {
	expresionA := &ExpresionTerminal{"Perro"}
	expresionB := &ExpresionTerminal{"Gato"}
	expresionC := &ExpresionTerminal{"Perro"}

	contextoOR := Contexto{"Perro"}
	expresionOR := &ExpresionOR{expresionA, expresionB}
	fmt.Printf("La expresion OR contiene la palabra perro: %v\n", expresionOR.Interpretar(contextoOR))

	contextoAND := Contexto{"Perro"}
	expresionAND := &ExpresionAND{expresionA, expresionC}
	fmt.Printf("La expresion AND contiene dos palabras perro: %v\n", expresionAND.Interpretar(contextoAND))
}
