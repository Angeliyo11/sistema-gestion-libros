package main

import "fmt"

// Estructura Libro representa un libro en el sistema.
type Libro struct {
	ID         int    // Identificador único del libro
	Titulo     string // Título del libro
	Autor      string // Autor del libro
	Anio       int    // Año de publicación del libro
	Disponible bool   // Estado de disponibilidad del libro (true o false)
}

// Base de datos simulada (slice de libros)
var libros = []Libro{
	{ID: 1, Titulo: "El Principito", Autor: "Antoine de Saint-Exupéry", Anio: 1943, Disponible: true},
	{ID: 2, Titulo: "1984", Autor: "George Orwell", Anio: 1949, Disponible: true},
	{ID: 3, Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Anio: 1967, Disponible: true},
}

// actualizarLibro actualiza la disponibilidad de un libro basado en su ID.
func actualizarLibro(id int, disponible bool) {
	for i, libro := range libros {
		if libro.ID == id {
			libros[i].Disponible = disponible
			fmt.Printf("Libro actualizado exitosamente: %+v\n", libros[i]) // Detalles del libro actualizado
			return
		}
	}
	fmt.Println("Libro no encontrado. No se pudo actualizar.")
}

// eliminarLibro elimina un libro de la lista de libros basado en su ID.
func eliminarLibro(id int) {
	for i, libro := range libros {
		if libro.ID == id {
			libros = append(libros[:i], libros[i+1:]...)             // Remueve el libro de la lista
			fmt.Printf("Libro eliminado exitosamente: %+v\n", libro) // Muestra el libro eliminado
			return
		}
	}
	fmt.Println("Libro no encontrado. No se pudo eliminar.")
}

// mostrarLibros muestra todos los libros actuales en la lista.
func mostrarLibros() {
	fmt.Println("\nLista de libros en el sistema:")
	for _, libro := range libros {
		fmt.Printf("- %+v\n", libro)
	}
}

// Función principal del sistema de gestión de libros
func main() {
	fmt.Println("Sistema de Gestión de Libros Electrónicos\n")

	// Muestra todos los libros disponibles inicialmente
	fmt.Println("Mostrando todos los libros disponibles:")
	mostrarLibros()

	// Actualiza la disponibilidad del libro con ID 1
	fmt.Println("\nActualizando disponibilidad del libro con ID 1...")
	actualizarLibro(1, false)

	// Elimina el libro con ID 2
	fmt.Println("\nEliminando el libro con ID 2...")
	eliminarLibro(2)

	// Muestra los libros restantes después de las operaciones
	fmt.Println("\nMostrando libros restantes después de actualizar y eliminar:")
	mostrarLibros()
}
