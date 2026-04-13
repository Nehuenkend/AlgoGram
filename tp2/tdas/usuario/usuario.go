package tdas

import (
	TDAHeap "tdas/cola_prioridad"
	TDAPost "tp2/tdas/post"
)

type Usuario interface {
	// ObtenerNombre devuelve el nombre del usuario.
	ObtenerNombre() string

	// ObtenerId devuelve el identificador único del usuario.
	ObtenerId() int

	// ObtenerFeed devuelve la cola de prioridad que representa el feed del usuario.
	ObtenerFeed() TDAHeap.ColaPrioridad[TDAPost.Post]
}
