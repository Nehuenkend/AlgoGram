package tdas

import (
	TDADict "tdas/diccionario"
)

type Post interface {
	// ObtenerID devuelve el identificador único del post.
	ObtenerID() int

	// ObtenerAutor devuelve el nombre del autor del post.
	ObtenerAutor() string

	// ObtenerIdAutor devuelve el identificador del autor del post.
	ObtenerIdAutor() int

	// ObtenerContenido devuelve el contenido del post.
	ObtenerContenido() string

	// CantidadLikes devuelve la cantidad de likes que tiene el post.
	CantidadLikes() int

	// ObtenerDictUsuariosLikes devuelve un diccionario con los usuarios que han dado like al post.
	ObtenerDictUsuariosLikes() TDADict.Diccionario[string, bool]

	// ObtenerUsuariosLikes devuelve una lista con los nombres de los usuarios que han dado like al post.
	ObtenerUsuariosLikes() []string

	// Likear permite a un usuario dar like al post.
	Likear(nombreUsuario string)
}
