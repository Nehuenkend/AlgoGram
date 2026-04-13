package cola_prioridad

const (
	COLA_VACIA        = "La cola esta vacia"
	CAPACIDAD_INICIAL = 1
	CAPACIDAD_MINIMA  = 1
	CTE_REDIMENSION   = 2
	CTE_REDUCCION     = 4
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func padre(i int) int {
	return (i - 1) / 2
}

func hijo_izq(i int) int {
	return (i * 2) + 1
}

func hijo_der(i int) int {
	return (i * 2) + 2
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic(COLA_VACIA)
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func upHeap[T any](arr []T, cmp func(T, T) int, i int) {
	for i > 0 && cmp(arr[padre(i)], arr[i]) < 0 {
		arr[padre(i)], arr[i] = arr[i], arr[padre(i)]
		i = padre(i)
	}
}

func downHeap[T any](arr []T, cmp func(T, T) int, cant, i int) {
	for {
		izq := hijo_izq(i)
		der := hijo_der(i)
		mayor := i

		if izq < cant && cmp(arr[izq], arr[mayor]) > 0 {
			mayor = izq
		}

		if der < cant && cmp(arr[der], arr[mayor]) > 0 {
			mayor = der
		}

		if mayor == i {
			break
		}

		arr[i], arr[mayor] = arr[mayor], arr[i]
		i = mayor
	}
}

func (cola *colaConPrioridad[T]) redimensionar(nuevaCap int) {
	nuevo := make([]T, nuevaCap)
	copy(nuevo, cola.datos[:cola.cant])
	cola.datos = nuevo
}

func (cola *colaConPrioridad[T]) Encolar(t T) {
	if cola.cant == cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) * CTE_REDIMENSION)
	}

	cola.cant++
	cola.datos[cola.cant-1] = t
	upHeap(cola.datos, cola.cmp, cola.cant-1)
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic(COLA_VACIA)
	}
	dato := cola.datos[0]
	cola.datos[cola.cant-1], cola.datos[0] = cola.datos[0], cola.datos[cola.cant-1]
	cola.cant--
	downHeap(cola.datos, cola.cmp, cola.cant, 0)

	if cola.cant <= cap(cola.datos)/CTE_REDUCCION && cap(cola.datos) > CAPACIDAD_MINIMA {
		nuevaCap := cap(cola.datos) / CTE_REDUCCION
		if nuevaCap < CAPACIDAD_MINIMA {
			nuevaCap = CAPACIDAD_MINIMA
		}

		cola.redimensionar(nuevaCap)
	}

	return dato
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, CAPACIDAD_INICIAL),
		cant:  0,
		cmp:   cmp,
	}
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		downHeap(arr, cmp, len(arr), i)
	}
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	capInit := len(arreglo)
	if capInit < CAPACIDAD_INICIAL {
		capInit = CAPACIDAD_INICIAL
	}

	datos := make([]T, capInit)
	copy(datos, arreglo)
	heap := colaConPrioridad[T]{
		datos: datos,
		cant:  len(arreglo),
		cmp:   cmp,
	}

	heapify(heap.datos, cmp)
	return &heap
}

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	heapify(elementos, cmp)

	for limite := len(elementos) - 1; limite > 0; limite-- {
		elementos[0], elementos[limite] = elementos[limite], elementos[0]
		downHeap(elementos, cmp, limite, 0)
	}
}
