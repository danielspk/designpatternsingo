package main

import (
	"fmt"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/chainofresponsability"
	"github.com/danielspk/designpatternsingo/patrones/comportamiento/strategy"
)

const LineaDivisora = "###########################################################"
const LineaSubDivisora = "-----------------------------------------------------------"

func main() {
	fmt.Println(LineaDivisora + "\nEjecutando Patrones de Comportamiento\n" + LineaDivisora)

	runPatronStrategy()
	runPatronChainOfResponsability()
}

func runPatronStrategy() {
	fmt.Println("\n" + LineaSubDivisora + "\nPatrón Strategy\n" + LineaSubDivisora + "\n")

	strategyContexto := strategy.Contexto{}
	fmt.Println(strategyContexto)
}

func runPatronChainOfResponsability() {
	fmt.Println("\n" + LineaSubDivisora + "\nPatrón Chain of Responsability\n" + LineaSubDivisora + "\n")

	chainReceptor := chainofresponsability.ReceptorAltaPrioridad{}
	fmt.Println(chainReceptor)
}
