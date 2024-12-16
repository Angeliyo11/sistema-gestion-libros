package main

import (
	"encoding/json" // Permite codificar y decodificar JSON
	"fmt"           // Proporciona funciones de entrada/salida
	"html/template" // Permite renderizar plantillas HTML
	"net/http"      // Facilita la creación de un servidor web
	"strconv"       // Convierte cadenas a tipos numéricos
)

// Estructura Libro: Representa un libro en el sistema
// Los tags `json` permiten generar o leer JSON con estos nombres

type Libro struct {
	ID         int    `json:"id"`         // Identificador del libro
	Titulo     string `json:"titulo"`     // Título del libro
	Autor      string `json:"autor"`      // Autor del libro
	Anio       int    `json:"anio"`       // Año de publicación
	Disponible bool   `json:"disponible"` // Disponibilidad del libro
}

// Base de datos simulada: Slice con información inicial de libros
var libros = []Libro{
	{ID: 1, Titulo: "El Principito", Autor: "Antoine de Saint-Exupéry", Anio: 1943, Disponible: true},
	{ID: 2, Titulo: "1984", Autor: "George Orwell", Anio: 1949, Disponible: true},
	{ID: 3, Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Anio: 1967, Disponible: true},
	{ID: 4, Titulo: "Don Quijote de la Mancha", Autor: "Miguel de Cervantes", Anio: 1605, Disponible: true},
	{ID: 5, Titulo: "La Casa de los Espíritus", Autor: "Isabel Allende", Anio: 1982, Disponible: true},
	{ID: 6, Titulo: "El Aleph", Autor: "Jorge Luis Borges", Anio: 1945, Disponible: true},
	{ID: 7, Titulo: "Crónica de una Muerte Anunciada", Autor: "Gabriel García Márquez", Anio: 1981, Disponible: true},
	{ID: 8, Titulo: "Ficciones", Autor: "Jorge Luis Borges", Anio: 1944, Disponible: true},
	{ID: 9, Titulo: "Rayuela", Autor: "Julio Cortázar", Anio: 1963, Disponible: true},
	{ID: 10, Titulo: "La Sombra del Viento", Autor: "Carlos Ruiz Zafón", Anio: 2001, Disponible: true},
}

// Función para mostrar todos los libros usando una plantilla HTML
func mostrarLibros(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html")) // Carga la plantilla HTML
	tmpl.Execute(w, libros)                                  // Renderiza la plantilla con la lista de libros
}

// Función para buscar un libro por su ID en formato JSON
func buscarLibro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Obtiene el ID desde la URL
	var libroEncontrado *Libro    // Puntero para almacenar el libro encontrado

	// Itera sobre la lista de libros para buscar coincidencias
	for _, libro := range libros {
		if fmt.Sprintf("%d", libro.ID) == id { // Convierte ID a cadena y compara
			libroEncontrado = &libro
			break
		}
	}

	// Si se encuentra, se responde en formato JSON; si no, error
	if libroEncontrado != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libroEncontrado)
	} else {
		http.Error(w, "Libro no encontrado", http.StatusNotFound)
	}
}

// Función para actualizar el estado de disponibilidad de un libro
func actualizarLibro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")                           // Obtiene el ID desde la URL
	disponible := r.URL.Query().Get("disponible") == "true" // Verifica si es true o false
	var libroActualizado *Libro                             // Puntero para almacenar el libro actualizado

	// Itera sobre la lista de libros para actualizar su estado
	for i, libro := range libros {
		if fmt.Sprintf("%d", libro.ID) == id {
			libros[i].Disponible = disponible
			libroActualizado = &libros[i]
			break
		}
	}

	// Responde en JSON el libro actualizado o devuelve un error
	if libroActualizado != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libroActualizado)
	} else {
		http.Error(w, "Libro no encontrado", http.StatusNotFound)
	}
}

// Función para eliminar un libro por su ID
func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") // Obtiene el ID del libro a eliminar

	// Itera y elimina el libro encontrado
	for i, libro := range libros {
		if fmt.Sprintf("%d", libro.ID) == id {
			libros = append(libros[:i], libros[i+1:]...) // Elimina el libro del slice
			w.Write([]byte("Libro eliminado exitosamente"))
			return
		}
	}
	http.Error(w, "Libro no encontrado", http.StatusNotFound)
}

// Función para agregar un nuevo libro
func agregarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { // Solo permite solicitudes POST
		titulo := r.FormValue("titulo")
		autor := r.FormValue("autor")
		anio, _ := strconv.Atoi(r.FormValue("anio"))
		disponible := r.FormValue("disponible") == "on"

		// Crea un nuevo libro con los datos proporcionados
		libro := Libro{
			ID:         len(libros) + 1,
			Titulo:     titulo,
			Autor:      autor,
			Anio:       anio,
			Disponible: disponible,
		}
		libros = append(libros, libro)                // Agrega el libro a la lista
		http.Redirect(w, r, "/", http.StatusSeeOther) // Redirige a la página principal
	}
}

// Función principal para manejar rutas y lanzar el servidor
func main() {
	http.HandleFunc("/", mostrarLibros)             // Página principal
	http.HandleFunc("/buscar", buscarLibro)         // Buscar libro por ID
	http.HandleFunc("/actualizar", actualizarLibro) // Actualizar disponibilidad
	http.HandleFunc("/eliminar", eliminarLibro)     // Eliminar libro
	http.HandleFunc("/agregar", agregarLibro)       // Agregar un libro nuevo

	fmt.Println("Servidor iniciado en :8080")
	http.ListenAndServe(":8080", nil) // Inicia el servidor en el puerto 8080
}
