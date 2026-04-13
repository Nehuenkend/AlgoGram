package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estado int

const (
	VACIA estado = iota
	OCUPADA
	BORRADA
	CAPACIDAD_INICIAL   = 1
	CAPACIDAD_MINIMA    = 1
	FACTOR_MAXIMO       = 0.7
	FACTOR_MINIMO       = 0.25
	CTE_REDIMENSION     = 2
	ERROR_NO_PERTENECE  = "La clave no pertenece al diccionario"
	ITERADOR_FINALIZADO = "El iterador termino de iterar"
)

type celdaHash[K, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K, V any] struct {
	tabla     []celdaHash[K, V]
	cmp       func(K, K) bool
	capacidad int
	cantidad  int
	borrados  int
}

// FUNCION DE HASHING

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// FNV Hasher
func calculateFNVHash(s []byte) uint64 {
	h := fnv.New64a() // FNV-1a 64-bit
	h.Write(s)
	return h.Sum64()
}

// FUNCIONES AUXILIARES HASH CERRADO

func (hash *hashCerrado[K, V]) crearNuevaTabla(nuevaCap int) {
	hash.tabla = make([]celdaHash[K, V], nuevaCap)
	hash.capacidad = nuevaCap
	hash.cantidad = 0
	hash.borrados = 0
}

func (hash *hashCerrado[K, V]) requiereRedimension() int {
	// Checkeo FACTOR_MAXIMO
	factorDeCarga := float64(hash.cantidad+hash.borrados) / float64(hash.capacidad)
	if factorDeCarga >= FACTOR_MAXIMO {
		return hash.capacidad * CTE_REDIMENSION
	}
	// Checkeo FACTOR_MINIMO
	factorDeCarga = float64(hash.cantidad) / float64(hash.capacidad)
	if factorDeCarga <= FACTOR_MINIMO {
		nuevaCap := hash.capacidad / CTE_REDIMENSION
		if nuevaCap < hash.cantidad {
			nuevaCap = hash.cantidad
		}
		if nuevaCap < CAPACIDAD_INICIAL {
			nuevaCap = CAPACIDAD_INICIAL
		}
		return nuevaCap
	}
	return -1
}

func (hash *hashCerrado[K, V]) redimensionar(nuevaCap int) {
	vieja := hash.tabla
	hash.crearNuevaTabla(nuevaCap)
	for _, celda := range vieja {
		if celda.estado == OCUPADA {
			hash.guardar(celda.clave, celda.dato)
		}
	}
}

// busca el indice en el que iria (o esta) la clave pasada por parametro, -1 si no encuentra

func (hash *hashCerrado[K, V]) buscarIndice(clave K) (int, estado) {
	posIni := int(calculateFNVHash(convertirABytes(clave))) & (hash.capacidad - 1)
	actual := posIni
	primerBorrado := -1

	for {
		celda := &hash.tabla[actual]
		switch celda.estado {
		case VACIA:
			if primerBorrado != -1 {
				return primerBorrado, BORRADA
			}
			return actual, VACIA
		case OCUPADA:
			if hash.cmp(celda.clave, clave) {
				return actual, OCUPADA
			}
		case BORRADA:
			if primerBorrado == -1 {
				primerBorrado = actual
			}
		}

		actual = (actual + 1) % hash.capacidad
		if actual == posIni {
			break
		}
	}
	return primerBorrado, VACIA
}

func (hash *hashCerrado[K, V]) guardar(clave K, dato V) {
	pos, estado := hash.buscarIndice(clave)

	if estado == OCUPADA {
		hash.tabla[pos].dato = dato
	} else {
		hash.tabla[pos] = celdaHash[K, V]{clave, dato, OCUPADA}
		hash.cantidad++
		if estado == BORRADA {
			hash.borrados--
		}
	}
}

// HASH CERRADO

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	redim := hash.requiereRedimension()

	if redim != -1 {
		hash.redimensionar(redim)
	}

	hash.guardar(clave, dato)
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, estado := hash.buscarIndice(clave)
	return estado == OCUPADA
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos, estado := hash.buscarIndice(clave)

	if estado != OCUPADA {
		panic(ERROR_NO_PERTENECE)
	}

	return hash.tabla[pos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos, estado := hash.buscarIndice(clave)

	if estado != OCUPADA {
		panic(ERROR_NO_PERTENECE)
	}

	dato := hash.tabla[pos].dato
	hash.tabla[pos].estado = BORRADA
	hash.borrados++
	hash.cantidad--

	redim := hash.requiereRedimension()

	if redim != -1 {
		hash.redimensionar(redim)
	}

	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < int(hash.capacidad); i++ {
		celda := &hash.tabla[i]
		// busco una celda ocupada
		if celda.estado == OCUPADA {
			// si visitar devuelve false, se corta la iteración
			if !visitar(celda.clave, celda.dato) {
				return
			}
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashCerrado[K, V]{
		hash:     hash,
		posicion: -1,
	}
	iter.avanzar()
	return iter
}

func CrearHash[K, V any](cmp func(K, K) bool) Diccionario[K, V] {
	hash := &hashCerrado[K, V]{}
	hash.crearNuevaTabla(CAPACIDAD_INICIAL)
	hash.cmp = cmp
	return hash
}

// ITERADOR EXTERNO

type iterHashCerrado[K, V any] struct {
	hash     *hashCerrado[K, V]
	posicion int
}

// avanzar mueve el iterador a la siguiente posición ocupada
func (iter *iterHashCerrado[K, V]) avanzar() {
	iter.posicion++
	for iter.HaySiguiente() && iter.hash.tabla[iter.posicion].estado != OCUPADA {
		iter.posicion++
	}
}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iter.posicion < iter.hash.capacidad
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(ITERADOR_FINALIZADO)
	}
	return iter.hash.tabla[iter.posicion].clave, iter.hash.tabla[iter.posicion].dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(ITERADOR_FINALIZADO)
	}
	iter.avanzar()
}
