package main

import "fmt"

// Definimos la estructura Libro, que representa los detalles de un libro en el sistema.
// Un libro tiene un ID, título, autor, año de publicación y disponibilidad.
type Libro struct {
	ID         int    // Identificador único del libro
	Titulo     string // Título del libro
	Autor      string // Autor del libro
	Anio       int    // Año de publicación del libro
	Disponible bool   // Estado de disponibilidad del libro (true o false)
}

// Lista de libros simulada como base de datos
// En un sistema real, estos datos se almacenarían en una base de datos.
var libros = []Libro{
	{ID: 1, Titulo: "El Principito", Autor: "Antoine de Saint-Exupéry", Anio: 1943, Disponible: true},
	{ID: 2, Titulo: "1984", Autor: "George Orwell", Anio: 1949, Disponible: true},
}

// Función actualizarLibro permite actualizar las propiedades de un libro en la lista de libros.
// Parametros:
// - id: el identificador único del libro que se quiere actualizar.
// - disponible: el nuevo estado de disponibilidad del libro (true o false).
// Si el libro con el ID especificado es encontrado, su propiedad 'Disponible' se actualiza.
func actualizarLibro(id int, disponible bool) {
	for i, libro := range libros {
		if libro.ID == id {
			libros[i].Disponible = disponible              // Actualiza la propiedad 'Disponible'
			fmt.Println("Libro actualizado exitosamente.") // Mensaje de éxito
			return
		}
	}
	fmt.Println("Libro no encontrado.") // Mensaje si el libro no es encontrado
}

// Función eliminarLibro permite eliminar un libro de la lista de libros mediante su ID.
// Parametro:
// - id: el identificador único del libro que se quiere eliminar.
// Si el libro con el ID especificado es encontrado, se elimina de la lista de libros.
func eliminarLibro(id int) {
	for i, libro := range libros {
		if libro.ID == id {
			libros = append(libros[:i], libros[i+1:]...) // Elimina el libro de la lista
			fmt.Println("Libro eliminado exitosamente.") // Mensaje de éxito
			return
		}
	}
	fmt.Println("Libro no encontrado.") // Mensaje si el libro no es encontrado
}

// Función principal (main) que simula las operaciones del sistema de gestión de libros.
func main() {
	// Muestra un mensaje inicial para indicar el inicio del sistema.
	fmt.Println("Sistema de Gestión de Libros Electrónicos")

	// Mostramos el primer libro en la lista
	libro := libros[0]
	fmt.Println("Libro encontrado:", libro)

	// Actualizamos la disponibilidad del libro con ID 1 a 'false' (no disponible)
	actualizarLibro(1, false)

	// Eliminamos el libro con ID 1
	eliminarLibro(1)

	// Intentamos mostrar el libro restante
	// Esto muestra el libro restante en la lista después de la actualización y eliminación.
	libro = libros[0] // Actualiza la referencia al primer libro en la lista
	fmt.Println("Libro restante:", libro)
}
