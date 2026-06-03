package tdas

import (
	TDAPost "main/src/post"
	TDAHeap "tdas/cola_prioridad"
)

type Usuario interface {
	// ObtenerNombre devuelve el nombre del usuario.
	ObtenerNombre() string

	// ObtenerId devuelve el identificador único del usuario.
	ObtenerId() int

	// ObtenerFeed devuelve la cola de prioridad que representa el feed del usuario.
	ObtenerFeed() TDAHeap.ColaPrioridad[TDAPost.Post]
}
