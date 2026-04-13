package tdas

import (
	TDAHeap "tdas/cola_prioridad"
	TDAPost "tp2/tdas/post"
	"tp2/utils"
)

type usuario struct {
	nombre string
	feed   TDAHeap.ColaPrioridad[TDAPost.Post]
	id     int
}

func CrearUsuario(nombre string, id int) Usuario {
	feed := TDAHeap.CrearHeap(func(p1, p2 TDAPost.Post) int {
		// Heap de minimos por Afinidad
		afinidad := utils.Modulo(p2.ObtenerIdAutor()-id) - utils.Modulo(p1.ObtenerIdAutor()-id)

		if afinidad != 0 {
			return afinidad
		}

		// Si tienen la misma afinidad, es por Id
		return p2.ObtenerID() - p1.ObtenerID()
	})

	return &usuario{
		nombre: nombre,
		feed:   feed,
		id:     id,
	}
}

func (u *usuario) ObtenerNombre() string {
	return u.nombre
}

func (u *usuario) ObtenerId() int {
	return u.id
}

func (u *usuario) ObtenerFeed() TDAHeap.ColaPrioridad[TDAPost.Post] {
	return u.feed
}
