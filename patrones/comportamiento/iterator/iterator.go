package iterator

import "fmt"

// Iterador Interface
type Iterador interface {
	Valor() string
	Siguiente()
	Anterior()
}

// Agregado Interface
type Agregado interface {
	CrearIterador() Iterador
}

// Agregado Concreto
type Radio struct {
	Estaciones []string
}

func (r *Radio) CrearIterador() Iterador {
	return &RadioIterador{radio: r}
}

func (r *Radio) Registrar(estacion string) {
	r.Estaciones = append(r.Estaciones, estacion)
}

// Iterador Concreto
type RadioIterador struct {
	radio  *Radio
	indice int
}

func (ri *RadioIterador) Valor() string {
	return ri.radio.Estaciones[ri.indice]
}

func (ri *RadioIterador) Siguiente() {
	ri.indice++
}

func (ri *RadioIterador) Anterior() {
	ri.indice--
}

// Test de uso del patrón
func TestPattern() {
	radio := &Radio{}
	radio.Registrar("FM100")
	radio.Registrar("FM200")
	radio.Registrar("FM300")

	iterador := radio.CrearIterador()

	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Anterior()
	fmt.Printf("Escuhando nuevamente la radio %s\n", iterador.Valor())
}
