package usuario

import (
	TDAPost "algogram/main/post"
	utils "algogram/main/utils"
	TDAHeap "tdas/cola_prioridad"
)

type usuario struct {
	nombre string
	feed   TDAHeap.ColaPrioridad[TDAPost.Post]
	id     int
}

func CrearUsuario(nombre string, id int) Usuario {
	feed := TDAHeap.CrearHeap(func(p1, p2 TDAPost.Post) int {
		// Heap de mínimos ordenado por afinidad (distancia del autor respecto al usuario actual)
		// Luego por ID del post (para determinismo)
		distancia1 := utils.Modulo(p1.ObtenerIdAutor() - id)
		distancia2 := utils.Modulo(p2.ObtenerIdAutor() - id)
		distanciaModulo := distancia2 - distancia1

		if distanciaModulo != 0 {
			return distanciaModulo
		}

		// Si tienen la misma distancia, ordenar por ID del post (descendente)
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
