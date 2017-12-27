package strategy

// Interface
type Estrategia interface {
	RealizarOperacion(int, int) int
}

// Estrategia Suma
type EstrategiaSuma struct{}

func (e EstrategiaSuma) RealizarOperacion(num1 int, num2 int) int {
	return num1 + num2
}

// Estrategia Resta
type EstrategiaResta struct{}

func (e EstrategiaResta) RealizarOperacion(num1 int, num2 int) int {
	return num1 - num2
}

// Estrategia Multiplica
type EstrategiaMultiplica struct{}

func (e EstrategiaMultiplica) RealizarOperacion(num1 int, num2 int) int {
	return num1 * num2
}

// Contexto
type Contexto struct {
	estrategia Estrategia
}

func (c *Contexto) EjecutarOperacion(num1 int, num2 int) int {
	return c.estrategia.RealizarOperacion(num1, num2)
}
