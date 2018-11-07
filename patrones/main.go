package main

import (
	"fmt"

	"github.com/danielspk/designpatternsingo/patrones/comportamiento/chainofresponsability"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/command"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/interpreter"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/iterator"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/mediator"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/memento"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/observer"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/state"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/strategy"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/templatemethod"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/visitor"
	"github.com/danielspk/designpatternsingo/patrones/creacionales/singleton"
)

const LineaDivisora = "################################################################"
const LineaSubDivisora = "----------------------------------------------------------------"

func main() {
	fmt.Printf("%s\nEjecutando Patrones de Comportamiento\n%s\n", LineaDivisora, LineaDivisora)

	fmt.Printf("\n%s\nPatrón Strategy\n%s\n", LineaSubDivisora, LineaSubDivisora)
	strategy.TestPattern()

	fmt.Printf("\n%s\nPatrón Chain of Responsability\n%s\n", LineaSubDivisora, LineaSubDivisora)
	chainofresponsability.TestPattern()

	fmt.Printf("\n%s\nPatrón Command\n%s\n", LineaSubDivisora, LineaSubDivisora)
	command.TestPattern()

	fmt.Printf("\n%s\nPatrón Template Method\n%s\n", LineaSubDivisora, LineaSubDivisora)
	templatemethod.TestPattern()

	fmt.Printf("\n%s\nPatrón Memento\n%s\n", LineaSubDivisora, LineaSubDivisora)
	memento.TestPattern()

	fmt.Printf("\n%s\nPatrón Interpreter\n%s\n", LineaSubDivisora, LineaSubDivisora)
	interpreter.TestPattern()

	fmt.Printf("\n%s\nPatrón Iterator\n%s\n", LineaSubDivisora, LineaSubDivisora)
	iterator.TestPattern()

	fmt.Printf("\n%s\nPatrón Visitor\n%s\n", LineaSubDivisora, LineaSubDivisora)
	visitor.TestPattern()

	fmt.Printf("\n%s\nPatrón State\n%s\n", LineaSubDivisora, LineaSubDivisora)
	state.TestPattern()

	fmt.Printf("\n%s\nPatrón Mediator\n%s\n", LineaSubDivisora, LineaSubDivisora)
	mediator.TestPattern()

	fmt.Printf("\n%s\nPatrón Observer\n%s\n", LineaSubDivisora, LineaSubDivisora)
	observer.TestPattern()

	fmt.Printf("\n%s\nEjecutando Patrones Creacionales\n%s\n", LineaDivisora, LineaDivisora)

	fmt.Printf("\n%s\nPatrón Singleton\n%s\n", LineaSubDivisora, LineaSubDivisora)
	singleton.TestPattern()
}
