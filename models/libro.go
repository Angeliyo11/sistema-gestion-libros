package models

// Paquete models contiene las estructuras principales del sistema.

// Libro define el modelo de datos para representar un libro en el sistema.
type Libro struct {
	ID         string // Identificador único del libro
	Titulo     string // Título del libro
	Autor      string // Autor del libro
	Anio       int    // Año de publicación
	Disponible bool   // Indica si el libro está disponible para préstamo o consulta
}
