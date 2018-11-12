package decorator

import "fmt"

// Componente Interface
type CafeInterface interface {
	GetCosto() int
	GetDetalle() string
}

// Componente Concreto
type Cafe struct{}

func (c *Cafe) GetCosto() int {
	return 30
}

func (c *Cafe) GetDetalle() string {
	return "Cafe"
}

// Decorador
type CafeDecorador struct {
	CafeInterface
}

func (cd *CafeDecorador) GetCosto() int {
	return cd.CafeInterface.GetCosto()
}

func (cd *CafeDecorador) GetDetalle() string {
	return cd.CafeInterface.GetDetalle()
}

// Decorador Concreto
type CafeConCrema struct {
	*CafeDecorador
}

func (ccc *CafeConCrema) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 15
}

func (ccc *CafeConCrema) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con crema"
}

// Decorador Concreto
type CafeConCanela struct {
	*CafeDecorador
}

func (ccc *CafeConCanela) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 10
}

func (ccc *CafeConCanela) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con canela"
}

// Test de uso del patrón
func TestPattern() {
	cafe := &Cafe{}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafe.GetDetalle(), cafe.GetCosto())

	cafeConCrema := &CafeConCrema{&CafeDecorador{cafe}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCrema.GetDetalle(), cafeConCrema.GetCosto())

	cafeConCremaConCanela := &CafeConCanela{&CafeDecorador{cafeConCrema}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCremaConCanela.GetDetalle(), cafeConCremaConCanela.GetCosto())
}
