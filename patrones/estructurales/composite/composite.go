package composite

import "fmt"

// Componente Interface
type Empleado interface {
	ObtenerSalario() int
}

// Hoja
type DesarrolladorSenior struct{}

func (ds *DesarrolladorSenior) ObtenerSalario() int {
	return 1000
}

// Hoja
type DesarrolladorJunior struct{}

func (dj *DesarrolladorJunior) ObtenerSalario() int {
	return 750
}

// Compuesto
type GerenciaIT struct {
	empleados []Empleado
}

func (g *GerenciaIT) AgregarEmpleado(empleado Empleado) {
	g.empleados = append(g.empleados, empleado)
}

func (g *GerenciaIT) ObtenerSalario() int {
	sumaSalarios := 0

	for _, empleado := range g.empleados {
		sumaSalarios = sumaSalarios + empleado.ObtenerSalario()
	}

	return sumaSalarios
}

// Test de uso del patrón
func TestPattern() {
	empleadoA := &DesarrolladorJunior{}
	empleadoB := &DesarrolladorJunior{}
	empleadoC := &DesarrolladorSenior{}

	gerenciaIT := &GerenciaIT{}
	gerenciaIT.AgregarEmpleado(empleadoA)
	gerenciaIT.AgregarEmpleado(empleadoB)
	gerenciaIT.AgregarEmpleado(empleadoC)

	fmt.Printf("El salario individual del desarrollador A es de $%d\n", empleadoA.ObtenerSalario())
	fmt.Printf("El salario individual del desarrollador B es de $%d\n", empleadoB.ObtenerSalario())
	fmt.Printf("El salario individual del desarrollador C es de $%d\n", empleadoC.ObtenerSalario())
	fmt.Printf("Los salarios de todos los desarrolladores de la Gerencia es de $%d\n", gerenciaIT.ObtenerSalario())
}
