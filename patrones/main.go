package main

import (
	"fmt"

	"github.com/danielspk/designpatternsingo/patrones/comportamiento/chainofresponsability"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/strategy"
)

const LineaDivisora = "###########################################################"
const LineaSubDivisora = "-----------------------------------------------------------"

func main() {
	fmt.Printf("%s\nEjecutando Patrones de Comportamiento\n%s\n", LineaDivisora, LineaDivisora)

	fmt.Printf("%s\nPatrón Strategy\n%s\n", LineaSubDivisora, LineaSubDivisora)
	strategy.TestPattern()

	fmt.Printf("%s\nPatrón Chain of Responsability\n%s\n", LineaSubDivisora, LineaSubDivisora)
	chainofresponsability.TestPattern()
}
