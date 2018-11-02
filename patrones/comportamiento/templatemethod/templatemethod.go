package templatemethod

import "fmt"

// Clase Abstracta - Interface
type DeployInterface interface {
	Testear()
	Compilar()
	Publicar()
}

// Clase Abstracta
type Deploy struct{}

// Método Plantilla
func (d Deploy) Construir(di DeployInterface) {
	fmt.Println("Ejecutando las siguientes acciones:")

	di.Testear()
	di.Compilar()
	di.Publicar()
}

// Clase Concreta - Android
type DeployAndroid struct {
	Deploy
}

func (d DeployAndroid) Testear() {
	fmt.Println("Android: Testeando")
}

func (d DeployAndroid) Compilar() {
	fmt.Println("Android: Compilando")
}

func (d DeployAndroid) Publicar() {
	fmt.Println("Android: Publicando")
}

// Clase Concreta - iOS
type DeployiOS struct {
	Deploy
}

func (d DeployiOS) Testear() {
	fmt.Println("iOS: Testeando")
}

func (d DeployiOS) Compilar() {
	fmt.Println("iOS: Compilando")
}

func (d DeployiOS) Publicar() {
	fmt.Println("iOS: Publicando")
}

// Test de uso del patrón
func TestPattern() {
	deployAndroid := DeployAndroid{Deploy{}}
	deployAndroid.Construir(&deployAndroid)

	fmt.Print("\n")

	deployiOS := DeployiOS{Deploy{}}
	deployiOS.Construir(&deployiOS)
}
