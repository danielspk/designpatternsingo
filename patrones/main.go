package main

import (
	"fmt"

	"github.com/danielspk/designpatternsingo/patrones/comportamiento/chainofresponsability"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/command"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/strategy"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/templatemethod"
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
}
