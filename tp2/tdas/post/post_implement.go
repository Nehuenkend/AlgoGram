package tdas

import (
	TDAHeap "tdas/cola_prioridad"
	TDADict "tdas/diccionario"
	"tp2/utils"
)

var _id int = -1

type post struct {
	id                int
	contenido         string
	likes             int
	autor             string
	id_autor          int
	usuariosLikes     TDADict.Diccionario[string, bool]
	heapUsuariosLikes TDAHeap.ColaPrioridad[string]
}

func CrearPost(autor, contenido string, autor_id int) Post {
	_id++
	return &post{
		id:                _id,
		contenido:         contenido,
		likes:             0,
		autor:             autor,
		id_autor:          autor_id,
		usuariosLikes:     TDADict.CrearHash[string, bool](func(a, b string) bool { return a == b }),
		heapUsuariosLikes: TDAHeap.CrearHeap(utils.StrCmp),
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

func (p *post) ObtenerDictUsuariosLikes() TDADict.Diccionario[string, bool] {
	return p.usuariosLikes
}

func (p *post) ObtenerUsuariosLikes() []string {
	backup := make([]string, 0)

	for !p.heapUsuariosLikes.EstaVacia() {
		nombreActual := p.heapUsuariosLikes.Desencolar()
		backup = append(backup, nombreActual)
	}

	p.heapUsuariosLikes = TDAHeap.CrearHeapArr(backup, utils.StrCmp)
	return backup
}

func (p *post) CantidadLikes() int {
	return p.likes
}

func (p *post) Likear(usuario string) {
	p.likes++
	p.usuariosLikes.Guardar(usuario, true)
	p.heapUsuariosLikes.Encolar(usuario)
}
