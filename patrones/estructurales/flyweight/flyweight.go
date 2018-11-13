package flyweight

import "fmt"

// Flyweight Interface
type Control interface {
	Dibujar(x string, y string) string
	GetReferencia() string
}

// Flyweight Concreto
type Boton struct {
	referencia string
}

func (b *Boton) Dibujar(x string, y string) string {
	return "Dibujando Botón #" + b.referencia + " en ejes " + x + ", " + y
}

func (b *Boton) GetReferencia() string {
	return b.referencia
}

// Flyweight Concreto No Compartido
type CampoPassword struct{}

func (cp *CampoPassword) Dibujar(x string, y string) string {
	return "Dibujando Campo de Password en ejes " + x + ", " + y
}

func (cp *CampoPassword) GetReferencia() string {
	return ""
}

// Fabrica Flyweight
type PantallaFabrica struct {
	controles []Control
}

func (pb *PantallaFabrica) ObtenerControl(tipo string, referencia string, x string, y string) string {
	for _, control := range pb.controles {
		if control.GetReferencia() == referencia {
			return control.Dibujar(x, y) + " || <control recuperado>"
		}
	}

	if tipo == "BOTON" {
		control := &Boton{referencia}

		pb.controles = append(pb.controles, control)

		return control.Dibujar(x, y)
	}

	if tipo == "PASSWORD" {
		control := &CampoPassword{}

		return control.Dibujar(x, y)
	}

	return ""
}

// Test de uso del patrón
func TestPattern() {
	pantalla := &PantallaFabrica{}

	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN1", "100", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN2", "200", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN3", "300", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("PASSWORD", "PWD1", "500", "300"))
	fmt.Printf("%s\n", pantalla.ObtenerControl("BOTON", "BTN1", "400", "300"))
}
