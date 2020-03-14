package main

import (
	"fmt"

	"designpatternsingo/patrones/comportamiento/chainofresponsibility"
	"designpatternsingo/patrones/comportamiento/command"
	"designpatternsingo/patrones/comportamiento/interpreter"
	"designpatternsingo/patrones/comportamiento/iterator"
	"designpatternsingo/patrones/comportamiento/mediator"
	"designpatternsingo/patrones/comportamiento/memento"
	"designpatternsingo/patrones/comportamiento/observer"
	"designpatternsingo/patrones/comportamiento/state"
	"designpatternsingo/patrones/comportamiento/strategy"
	"designpatternsingo/patrones/comportamiento/templatemethod"
	"designpatternsingo/patrones/comportamiento/visitor"
	"designpatternsingo/patrones/creacionales/abstractfactory"
	"designpatternsingo/patrones/creacionales/builder"
	"designpatternsingo/patrones/creacionales/factorymethod"
	"designpatternsingo/patrones/creacionales/prototype"
	"designpatternsingo/patrones/creacionales/singleton"
	"designpatternsingo/patrones/estructurales/adapter"
	"designpatternsingo/patrones/estructurales/bridge"
	"designpatternsingo/patrones/estructurales/composite"
	"designpatternsingo/patrones/estructurales/decorator"
	"designpatternsingo/patrones/estructurales/facade"
	"designpatternsingo/patrones/estructurales/flyweight"
	"designpatternsingo/patrones/estructurales/proxy"
)

const LineaDivisora = "################################################################"
const LineaSubDivisora = "----------------------------------------------------------------"

func main() {
	fmt.Printf("%s\nEjecutando Patrones de Comportamiento\n%s\n", LineaDivisora, LineaDivisora)

	fmt.Printf("\n%s\nPatrón Strategy\n%s\n", LineaSubDivisora, LineaSubDivisora)
	strategy.TestPattern()

	fmt.Printf("\n%s\nPatrón Chain of responsibility\n%s\n", LineaSubDivisora, LineaSubDivisora)
	chainofresponsibility.TestPattern()

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

	fmt.Printf("\n%s\nPatrón Builder\n%s\n", LineaSubDivisora, LineaSubDivisora)
	builder.TestPattern()

	fmt.Printf("\n%s\nPatrón Factory Method\n%s\n", LineaSubDivisora, LineaSubDivisora)
	factorymethod.TestPattern()

	fmt.Printf("\n%s\nPatrón Abstract Factory\n%s\n", LineaSubDivisora, LineaSubDivisora)
	abstractfactory.TestPattern()

	fmt.Printf("\n%s\nPatrón Prototype\n%s\n", LineaSubDivisora, LineaSubDivisora)
	prototype.TestPattern()

	fmt.Printf("\n%s\nEjecutando Patrones Estructurales\n%s\n", LineaDivisora, LineaDivisora)

	fmt.Printf("\n%s\nPatrón Composite\n%s\n", LineaSubDivisora, LineaSubDivisora)
	composite.TestPattern()

	fmt.Printf("\n%s\nPatrón Adapter\n%s\n", LineaSubDivisora, LineaSubDivisora)
	adapter.TestPattern()

	fmt.Printf("\n%s\nPatrón Bridge\n%s\n", LineaSubDivisora, LineaSubDivisora)
	bridge.TestPattern()

	fmt.Printf("\n%s\nPatrón Proxy\n%s\n", LineaSubDivisora, LineaSubDivisora)
	proxy.TestPattern()

	fmt.Printf("\n%s\nPatrón Decorator\n%s\n", LineaSubDivisora, LineaSubDivisora)
	decorator.TestPattern()

	fmt.Printf("\n%s\nPatrón Facade\n%s\n", LineaSubDivisora, LineaSubDivisora)
	facade.TestPattern()

	fmt.Printf("\n%s\nPatrón Flyweight\n%s\n", LineaSubDivisora, LineaSubDivisora)
	flyweight.TestPattern()
}
