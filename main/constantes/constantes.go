package constantes

const (
	COMANDO_LOGIN              = "login"
	COMANDO_LOGOUT             = "logout"
	COMANDO_PUBLICAR           = "publicar"
	COMANDO_VER_SIGUIENTE_FEED = "ver_siguiente_feed"
	COMANDO_LIKEAR_POST        = "likear_post"
	COMANDO_MOSTRAR_LIKES      = "mostrar_likes"

	COMANDO_NO_RECONOCIDO = "Comando no reconocido"
	BIENVENIDA            = "Hola"
	DESPEDIDA             = "Adios"
	POST_PUBLICADO        = "Post publicado"
	POST_LIKEADO          = "Post likeado"

	ERROR_USUARIO_NO_EXISTE = "Error: usuario no existente"
	ERROR_USUARIO_LOGIN     = "Error: Ya habia un usuario loggeado"
	ERROR_USUARIO_LOGOUT    = "Error: no habia usuario loggeado"
	ERROR_ARCHIVO           = "Error: no se pudo abrir el archivo"
	ERROR_NO_HAY_POSTS      = "Usuario no loggeado o no hay mas posts para ver"
	ERROR_LIKEAR_POST       = "Error: Usuario no loggeado o Post inexistente"
	ERROR_MOSTRAR_LIKES     = "Error: Post inexistente o sin likes"
	ERROR_ARCHIVO_STDIN     = "Error: No se proporciono el archivo de usuarios"
)
