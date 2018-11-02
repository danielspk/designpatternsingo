package memento

import "fmt"

// Interface
type Memento interface {
	SetContenido(string)
	GetContenido() string
}

// Memento
type EditorMemento struct {
	contenido string
}

func (em *EditorMemento) SetContenido(contenido string) {
	em.contenido = contenido
}

func (em *EditorMemento) GetContenido() string {
	return em.contenido
}

// Originador
type Editor struct {
	contenido string
}

func (e *Editor) VerContenido() string {
	return e.contenido
}

func (e *Editor) Escribir(texto string) {
	e.contenido = e.contenido + " " + texto
}

func (e *Editor) Guardar() Memento {
	editorMemento := &EditorMemento{}
	editorMemento.SetContenido(e.contenido)

	return editorMemento
}

func (e *Editor) Restaurar(memento Memento) {
	e.contenido = memento.GetContenido()
}

// Test de uso del patrón
func TestPattern() {
	editor := &Editor{}
	editor.Escribir("TextoA")
	editor.Escribir("TextoB")
	fmt.Printf("El editor contiene:%s\n", editor.VerContenido())

	fmt.Println("Se guarda el estado actual")
	memento := editor.Guardar()

	fmt.Println("Se escribe unnuevo contenido")
	editor.Escribir("TextoC")

	fmt.Printf("El editor ahora contiene:%s\n", editor.VerContenido())

	fmt.Println("Se restaura el contenido guardado")
	editor.Restaurar(memento)

	fmt.Printf("El editor nuevamente contiene:%s\n", editor.VerContenido())
}
