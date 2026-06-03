package post

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

	// ObtenerUsuariosLikes devuelve una copia de la lista de usuarios que han dado like al post.
	// Retorna una copia para evitar que el código cliente modifique la lista interna.
	ObtenerUsuariosLikes() []string

	// Likear permite a un usuario dar like al post.
	// Si el usuario ya ha likeado, no hace nada (idempotente).
	Likear(nombreUsuario string)

	// UsuarioYaLikeo verifica si un usuario ya ha dado like a este post.
	// Encapsula la lógica de búsqueda en el diccionario de likes.
	UsuarioYaLikeo(nombreUsuario string) bool
}
