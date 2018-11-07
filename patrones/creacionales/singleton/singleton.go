package singleton

import (
	"fmt"
	"sync"
	"time"
)

// Singleton
type Singleton struct {
	Tiempo int64
}

// Creador "estático"
var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			time.Now().Unix(),
		}
	})

	return instance
}

// Test de uso del patrón
func TestPattern() {
	fmt.Println("Todas las instancias Singleton tienen que tener el mismo número")

	fmt.Printf("Instancia Singleton: %d\n", GetInstance().Tiempo)

	time.Sleep(1 * time.Second)

	fmt.Printf("Instancia Singleton: %d\n", GetInstance().Tiempo)

	canalEspera := make(chan int64)

	go func() {
		time.Sleep(1 * time.Second)

		canalEspera <- GetInstance().Tiempo
	}()

	fmt.Printf("Instancia Singleton: %d\n", <-canalEspera)
}
