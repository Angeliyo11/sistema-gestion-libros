package models

// Paquete models contiene las estructuras principales del sistema, incluyendo la de usuarios.

// Usuario define el modelo de datos para representar un usuario en el sistema.
// Representa tanto a lectores como a administradores del sistema.
type Usuario struct {
	ID      string // Identificador único del usuario
	Nombre  string // Nombre completo del usuario
	Correo  string // Correo electrónico del usuario
	EsAdmin bool   // Indica si el usuario tiene privilegios de administrador
}
