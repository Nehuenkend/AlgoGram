package post

import (
	TDADict "tdas/diccionario"
)

type post struct {
	id               int
	contenido        string
	likes            int
	autor            string
	id_autor         int
	usuariosLikes    TDADict.Diccionario[string, bool]
	usuariosLikesArr []string // Mantener lista sincronizada para evitar side effects
}

// CrearPost crea un nuevo post con un ID generado por el generador inyectado
// Parámetros:
//   - generator: IDGenerator para obtener IDs únicos
//   - autor: nombre del usuario que publica
//   - contenido: texto del post
//   - autor_id: identificador numérico del autor
func CrearPost(generator IDGenerator, autor, contenido string, autor_id int) Post {
	return &post{
		id:               generator.Next(),
		contenido:        contenido,
		likes:            0,
		autor:            autor,
		id_autor:         autor_id,
		usuariosLikes:    TDADict.CrearHash[string, bool](func(a, b string) bool { return a == b }),
		usuariosLikesArr: make([]string, 0),
	}
}

func (p *post) ObtenerID() int {
	return p.id
}

func (p *post) ObtenerIdAutor() int {
	return p.id_autor
}

func (p *post) ObtenerAutor() string {
	return p.autor
}

func (p *post) ObtenerContenido() string {
	return p.contenido
}

// ObtenerUsuariosLikes retorna una copia de la lista de usuarios que likearon.
func (p *post) ObtenerUsuariosLikes() []string {
	copia := make([]string, len(p.usuariosLikesArr))
	copy(copia, p.usuariosLikesArr)
	return copia
}

func (p *post) CantidadLikes() int {
	return p.likes
}

// UsuarioYaLikeo encapsula la lógica de búsqueda en el diccionario.
func (p *post) UsuarioYaLikeo(nombreUsuario string) bool {
	return p.usuariosLikes.Pertenece(nombreUsuario)
}

func (p *post) Likear(usuario string) {
	if !p.usuariosLikes.Pertenece(usuario) {
		p.likes++
		p.usuariosLikes.Guardar(usuario, true)
		p.usuariosLikesArr = append(p.usuariosLikesArr, usuario)
	}
}
