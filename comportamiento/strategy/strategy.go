package strategy

// Interface
type Strategy interface {
	DoOperation(int, int) int
}

// Estrategia Add
type OpearationAdd struct{}

func (o OpearationAdd) DoOperation(num1 int, num2 int) int {
	return num1 + num2
}

// Estrategia Substract
type OpearationSubstract struct{}

func (o OpearationSubstract) DoOperation(num1 int, num2 int) int {
	return num1 - num2
}

// Estrategia Multiply
type OpearationMultiply struct{}

func (o OpearationMultiply) DoOperation(num1 int, num2 int) int {
	return num1 * num2
}

// Contexto
type Contect struct {
	strategy Strategy
}

func (c *Contect) ExecuteStrategy(num1 int, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}
