package command

import "fmt"

// Interface Comando (Orden)
type Comando interface {
	Ejecutar() string
}

// Comando Prender (OrdenConcreta)
type ComandoPrender struct {
	receptor Receptor
}

func (cp ComandoPrender) Ejecutar() string {
	return cp.receptor.Prender()
}

// Comando Apagar (OrdenConcreta)
type ComandoApagar struct {
	receptor Receptor
}

func (ca ComandoApagar) Ejecutar() string {
	return ca.receptor.Apagar()
}

// Invocador
type Invocador struct {
	comandos []Comando
}

func (i *Invocador) GuardarComando(comando Comando) {
	i.comandos = append(i.comandos, comando)
}

func (i *Invocador) EliminarUltimoComando() {
	if len(i.comandos) != 0 {
		i.comandos = i.comandos[:len(i.comandos)-1]
	}
}

func (i *Invocador) Limpiar() {
	i.comandos = []Comando{}
}

func (i *Invocador) Ejecutar() string {
	var resultados string

	for _, comando := range i.comandos {
		resultados += comando.Ejecutar() + "\n"
	}

	return resultados
}

// Receptor
type Receptor struct{}

func (r Receptor) Prender() string {
	return "- Prender Televisor"
}

func (r Receptor) Apagar() string {
	return "- Apagar Televisor"
}

// Test de uso del patrón
func TestPattern() {
	invocador := Invocador{comandos: []Comando{}}
	receptor := Receptor{}

	// se establecen dos comandos concretos y se los elimina
	// el invocador queda sin comandos que ejecutar
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.GuardarComando(ComandoApagar{receptor: receptor})
	invocador.Limpiar()

	// se establecen dos comandos concretos iguales y se elimina el último
	// el invocador queda con un único comando para ejecutar
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.GuardarComando(ComandoPrender{receptor: receptor})
	invocador.EliminarUltimoComando()

	// se establece un comando concreto más
	invocador.GuardarComando(ComandoApagar{receptor: receptor})

	// el invocador ejecuta los dos comandos
	fmt.Printf("Comandos ejecutados:\n%v\n", invocador.Ejecutar())
}
