package services

import "sistema-gestion-libros-electronicos/models"

// Paquete services contiene las funciones de manejo de datos de libros.

// Lista global de libros en memoria.
var Libros = []models.Libro{}

// AgregarLibro agrega un nuevo libro a la lista.
// Recibe como parámetro un objeto de tipo Libro.
func AgregarLibro(libro models.Libro) {
	Libros = append(Libros, libro)
}

// BuscarLibroPorID busca un libro en la lista basado en su ID.
// Devuelve un puntero al libro si lo encuentra, o nil si no existe.
func BuscarLibroPorID(id string) *models.Libro {
	for _, libro := range Libros {
		if libro.ID == id {
			return &libro
		}
	}
	return nil
}

// EliminarLibroPorID elimina un libro de la lista basado en su ID.
// Devuelve true si el libro fue eliminado, false si no se encontró.
func EliminarLibroPorID(id string) bool {
	for i, libro := range Libros {
		if libro.ID == id {
			// Eliminar el libro de la lista
			Libros = append(Libros[:i], Libros[i+1:]...)
			return true
		}
	}
	return false
}

// ActualizarLibro actualiza los datos de un libro existente en la lista.
// Devuelve true si el libro fue actualizado, false si no se encontró.
func ActualizarLibro(id string, nuevoLibro models.Libro) bool {
	for i, libro := range Libros {
		if libro.ID == id {
			// Actualizar los datos del libro
			Libros[i] = nuevoLibro
			return true
		}
	}
	return false
}
