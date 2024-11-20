package services

import "sistema-gestion-libros-electronicos/models"

// Paquete services contiene las funciones de manejo de datos de usuarios.

// Lista global para almacenar los usuarios del sistema en memoria.
var Usuarios = []models.Usuario{}

// AgregarUsuario agrega un nuevo usuario a la lista de usuarios.
// Recibe como parámetro un objeto de tipo Usuario.
func AgregarUsuario(usuario models.Usuario) {
	Usuarios = append(Usuarios, usuario)
}

// BuscarUsuarioPorID busca un usuario en la lista basado en su ID.
// Devuelve un puntero al usuario si lo encuentra, o nil si no existe.
func BuscarUsuarioPorID(id string) *models.Usuario {
	for _, usuario := range Usuarios {
		if usuario.ID == id {
			return &usuario
		}
	}
	return nil
}

// EliminarUsuarioPorID elimina un usuario de la lista basado en su ID.
// Devuelve true si el usuario fue eliminado, o false si no se encontró.
func EliminarUsuarioPorID(id string) bool {
	for i, usuario := range Usuarios {
		if usuario.ID == id {
			// Eliminar el usuario de la lista
			Usuarios = append(Usuarios[:i], Usuarios[i+1:]...)
			return true
		}
	}
	return false
}

// ActualizarUsuario actualiza los datos de un usuario existente en la lista.
// Recibe como parámetros el ID del usuario a actualizar y un nuevo objeto Usuario.
// Devuelve true si el usuario fue actualizado, o false si no se encontró.
func ActualizarUsuario(id string, nuevoUsuario models.Usuario) bool {
	for i, usuario := range Usuarios {
		if usuario.ID == id {
			// Actualizar los datos del usuario
			Usuarios[i] = nuevoUsuario
			return true
		}
	}
	return false
}
