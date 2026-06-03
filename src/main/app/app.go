package app

import (
	TDAUsuario "algogram/main/usuario"
	TDADiccionario "tdas/diccionario"
)

type App interface {
	// Login permite al usuario iniciar sesión con su nombre de usuario.
	Login(nombreUsuario string)

	// Logout permite al usuario cerrar sesión.
	Logout()

	// ObtenerUsuarioLogueado devuelve el usuario que ha iniciado sesión actualmente.
	ObtenerUsuarioLogueado() TDAUsuario.Usuario

	// CargarUsuarios carga los usuarios desde un archivo dado.
	// Retorna error si hay problemas al leer el archivo.
	CargarUsuarios(archivo string) error

	// ObtenerUsuariosRegistrados devuelve un diccionario con todos los usuarios registrados.
	ObtenerUsuariosRegistrados() TDADiccionario.Diccionario[string, TDAUsuario.Usuario]

	// Publicar permite al usuario publicar un contenido en su feed.
	Publicar(contenido string)

	// VerSiguienteFeed muestra el siguiente post en el feed del usuario.
	VerSiguienteFeed()

	// LikearPost permite al usuario dar like a un post específico.
	LikearPost(idPost int)

	// MostrarLikes muestra la cantidad de likes de un post específico.
	MostrarLikes(idPost int)
}
