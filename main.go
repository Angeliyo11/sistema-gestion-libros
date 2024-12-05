package main

import (
	"errors"
	"fmt"
)

// Libro representa los detalles de un libro en el sistema.
type Libro struct {
	id         int    // Identificador único del libro (privado)
	titulo     string // Título del libro (privado)
	autor      string // Autor del libro (privado)
	anio       int    // Año de publicación del libro (privado)
	disponible bool   // Estado de disponibilidad del libro (privado)
}

// Métodos encapsulados (Setters y Getters)
// SetID permite configurar el ID del libro de manera controlada.
func (l *Libro) SetID(id int) error {
	if id <= 0 {
		return errors.New("el ID debe ser un número positivo")
	}
	l.id = id
	return nil
}

// GetID devuelve el ID del libro.
func (l *Libro) GetID() int {
	return l.id
}

// SetDisponible permite modificar el estado de disponibilidad.
func (l *Libro) SetDisponible(disponible bool) {
	l.disponible = disponible
}

// GetDisponible devuelve el estado de disponibilidad del libro.
func (l *Libro) GetDisponible() bool {
	return l.disponible
}

// Operaciones con libros (Interfaz)
type OperacionesLibro interface {
	BuscarLibro(id int) (*Libro, error)
	ActualizarLibro(id int, disponible bool) error
	EliminarLibro(id int) error
}

// GestorLibros gestiona la lista de libros y sus operaciones.
type GestorLibros struct {
	libros []Libro
}

// BuscarLibro busca un libro por ID y lo devuelve.
func (g *GestorLibros) BuscarLibro(id int) (*Libro, error) {
	for _, libro := range g.libros {
		if libro.GetID() == id {
			return &libro, nil
		}
	}
	return nil, errors.New("libro no encontrado")
}

// ActualizarLibro actualiza el estado de disponibilidad de un libro.
func (g *GestorLibros) ActualizarLibro(id int, disponible bool) error {
	for i, libro := range g.libros {
		if libro.GetID() == id {
			g.libros[i].SetDisponible(disponible)
			fmt.Println("Libro actualizado exitosamente.")
			return nil
		}
	}
	return errors.New("libro no encontrado")
}

// EliminarLibro elimina un libro por su ID.
func (g *GestorLibros) EliminarLibro(id int) error {
	for i, libro := range g.libros {
		if libro.GetID() == id {
			g.libros = append(g.libros[:i], g.libros[i+1:]...)
			fmt.Println("Libro eliminado exitosamente.")
			return nil
		}
	}
	return errors.New("libro no encontrado")
}

// MostrarLibros imprime todos los libros disponibles en el sistema.
func (g *GestorLibros) MostrarLibros() {
	fmt.Println("Lista de libros:")
	for _, libro := range g.libros {
		fmt.Printf("- ID: %d, Título: %s, Autor: %s, Año: %d, Disponible: %t\n",
			libro.GetID(), libro.titulo, libro.autor, libro.anio, libro.GetDisponible())
	}
}

// Función principal
func main() {
	// Inicializamos los libros
	librosIniciales := []Libro{
		{id: 1, titulo: "El Principito", autor: "Antoine de Saint-Exupéry", anio: 1943, disponible: true},
		{id: 2, titulo: "1984", autor: "George Orwell", anio: 1949, disponible: true},
		{id: 3, titulo: "Cien Años de Soledad", autor: "Gabriel García Márquez", anio: 1967, disponible: true},
		{id: 4, titulo: "Don Quijote de la Mancha", autor: "Miguel de Cervantes", anio: 1605, disponible: true},
		{id: 5, titulo: "El Hobbit", autor: "J.R.R. Tolkien", anio: 1937, disponible: true},
		{id: 6, titulo: "Orgullo y Prejuicio", autor: "Jane Austen", anio: 1813, disponible: true},
		{id: 7, titulo: "Matar a un Ruiseñor", autor: "Harper Lee", anio: 1960, disponible: true},
		{id: 8, titulo: "La Metamorfosis", autor: "Franz Kafka", anio: 1915, disponible: true},
		{id: 9, titulo: "Ulises", autor: "James Joyce", anio: 1922, disponible: true},
		{id: 10, titulo: "El Gran Gatsby", autor: "F. Scott Fitzgerald", anio: 1925, disponible: true}}

	// Creamos el gestor de libros
	gestor := GestorLibros{libros: librosIniciales}

	// Mostrar libros iniciales
	fmt.Println("Sistema de Gestión de Libros Electrónicos")
	gestor.MostrarLibros()

	// Buscar un libro
	fmt.Println("\nBuscando el libro con ID 1:")
	libro, err := gestor.BuscarLibro(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Libro encontrado: %+v\n", libro)
	}

	// Actualizar disponibilidad
	fmt.Println("\nActualizando disponibilidad del libro con ID 1:")
	err = gestor.ActualizarLibro(1, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		gestor.MostrarLibros()
	}

	// Eliminar un libro
	fmt.Println("\nEliminando el libro con ID 1:")
	err = gestor.EliminarLibro(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		gestor.MostrarLibros()
	}
}
