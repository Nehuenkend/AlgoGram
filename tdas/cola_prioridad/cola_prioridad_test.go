package cola_prioridad_test

import (
	"strings"
	"testing"

	TDAColaPrioridad "tdas/cola_prioridad"

	"github.com/stretchr/testify/require"
)

// FUNCIONES AUXILIARES

func pruebaVolumen(t *testing.T, heap TDAColaPrioridad.ColaPrioridad[int], cantidad int) {
	for i := 0; i < cantidad; i++ {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i+1, heap.Cantidad())
	}

	for i := cantidad - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
		require.Equal(t, i, heap.Cantidad())
	}

	require.True(t, heap.EstaVacia())
}

func verificarOrden[T any](t *testing.T, heap TDAColaPrioridad.ColaPrioridad[T], esperado []T) {
	for i, v := range esperado {
		require.Equal(t, v, heap.VerMax())
		require.Equal(t, v, heap.Desencolar())
		require.Equal(t, len(esperado)-i-1, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
}

func verificarVacio[T any](t *testing.T, heap TDAColaPrioridad.ColaPrioridad[T]) {
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, TDAColaPrioridad.COLA_VACIA, func() {
		heap.VerMax()
	})
	require.PanicsWithValue(t, TDAColaPrioridad.COLA_VACIA, func() {
		heap.Desencolar()
	})
}

func verificarHeapSort[T any](t *testing.T, entrada []T, esperado []T, comparador func(a, b T) int) {
	TDAColaPrioridad.HeapSort(entrada, comparador)
	require.Equal(t, esperado, entrada)
}

// HEAP

func TestColaPrioridadEstaVacia(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(strings.Compare)
	verificarVacio(t, heap)

	heap.Encolar("A")
	verificarOrden(t, heap, []string{"A"})
	verificarVacio(t, heap)
}

func TestColaPrioridadEncolarDesencolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(strings.Compare)

	heap.Encolar("B")
	heap.Encolar("A")
	heap.Encolar("C")
	heap.Encolar("D")
	heap.Encolar("F")

	verificarOrden(t, heap, []string{"F", "D", "C", "B", "A"})
}

func TestColaPrioridadVerMax(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })
	verificarVacio(t, heap)

	for i := range 50 {
		heap.Encolar(i)
		require.Equal(t, i, heap.VerMax())
		require.False(t, heap.EstaVacia())
		require.Equal(t, i+1, heap.Cantidad())
	}

	for i := 49; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())

		if i == 0 {
			require.PanicsWithValue(t, TDAColaPrioridad.COLA_VACIA, func() {
				heap.VerMax()
			})
		} else {
			require.Equal(t, i-1, heap.VerMax())
		}
	}
}

func TestColaPrioridadPruebaVolumen(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	pruebaVolumen(t, heap, 1000)
	pruebaVolumen(t, heap, 10000)
	pruebaVolumen(t, heap, 100000)
}

func TestColaPrioridadDuplicadosSimple(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(strings.Compare)
	inputs := []string{"b", "b", "a", "b", "c", "c"}
	verificarVacio(t, heap)

	for _, v := range inputs {
		heap.Encolar(v)
	}

	verificarOrden(t, heap, []string{"c", "c", "b", "b", "b", "a"})
	verificarVacio(t, heap)
}

func TestColaPrioridadDuplicadosTodosIguales(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })
	verificarVacio(t, heap)

	const reps = 1000
	for i := 0; i < reps; i++ {
		heap.Encolar(42)
	}

	require.Equal(t, reps, heap.Cantidad())
	require.Equal(t, 42, heap.VerMax())

	for i := reps; i > 0; i-- {
		require.Equal(t, 42, heap.VerMax())
		require.Equal(t, 42, heap.Desencolar())
		require.Equal(t, i-1, heap.Cantidad())
	}

	verificarVacio(t, heap)
}

func TestColaPrioridadHeapRedimensionYReuso(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(func(a, b int) int { return a - b })

	const n = 5000
	for i := 0; i < n; i++ {
		heap.Encolar(i)
	}
	for i := 0; i < n; i++ {
		heap.Desencolar()
	}

	verificarVacio(t, heap)

	heap.Encolar(1)
	heap.Encolar(2)
	heap.Encolar(0)
	verificarOrden(t, heap, []int{2, 1, 0})
}

// HEAP ARRAY

func TestColaPrioridadHeapArrayCrear(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]string{"A", "B", "C", "D", "E"}, strings.Compare)
	verificarOrden(t, heap, []string{"E", "D", "C", "B", "A"})

	heap = TDAColaPrioridad.CrearHeapArr([]string{"Z", "Y", "X", "W", "V", "U", "T"}, strings.Compare)
	verificarOrden(t, heap, []string{"Z", "Y", "X", "W", "V", "U", "T"})

	heap = TDAColaPrioridad.CrearHeapArr([]string{"G", "R", "Y", "A"}, strings.Compare)
	verificarOrden(t, heap, []string{"Y", "R", "G", "A"})

	heapInts := TDAColaPrioridad.CrearHeapArr([]int{}, func(a, b int) int { return a - b })
	verificarVacio(t, heapInts)
}

func TestColaPrioridadHeapArrayEncolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]int{10, 20, 30}, func(a, b int) int { return a - b })

	heap.Encolar(25)
	heap.Encolar(5)
	heap.Encolar(40)

	verificarOrden(t, heap, []int{40, 30, 25, 20, 10, 5})
}

func TestColaPrioridadEncolarConHeapArrayVacio(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]int{}, func(a, b int) int { return a - b })
	verificarVacio(t, heap)

	heap.Encolar(15)
	heap.Encolar(10)
	heap.Encolar(20)

	verificarOrden(t, heap, []int{20, 15, 10})
}

func TestColaPrioridadHeapArrayCrearDesdeSliceConCapExtra(t *testing.T) {
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 2, 1, 3

	heap := TDAColaPrioridad.CrearHeapArr(s, func(a, b int) int { return a - b })
	verificarOrden(t, heap, []int{3, 2, 1})
	heap = TDAColaPrioridad.CrearHeapArr(s, func(a, b int) int { return a - b })

	heap.Encolar(5)
	heap.Encolar(0)
	verificarOrden(t, heap, []int{5, 3, 2, 1, 0})
}

func TestColaPrioridadHeapArrayCrearDesdeNilSliceYEncolar(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeapArr([]int(nil), func(a, b int) int { return a - b })
	verificarVacio(t, heap)

	heap.Encolar(7)
	verificarOrden(t, heap, []int{7})
}

// HEAPSORT

func TestColaPrioridadHeapSortEnteros(t *testing.T) {
	entrada := []int{5, 3, 8, 1, 2, 7, 4, 6}
	esperado := []int{1, 2, 3, 4, 5, 6, 7, 8}
	verificarHeapSort(t, entrada, esperado, func(a, b int) int { return a - b })

	entrada = []int{10, -1, 2, 0, 5, 3}
	esperado = []int{-1, 0, 2, 3, 5, 10}
	verificarHeapSort(t, entrada, esperado, func(a, b int) int { return a - b })

	entrada = []int{}
	esperado = []int{}
	verificarHeapSort(t, entrada, esperado, func(a, b int) int { return a - b })

	entrada = []int{42}
	esperado = []int{42}
	verificarHeapSort(t, entrada, esperado, func(a, b int) int { return a - b })

	entrada = []int{3, 3, 3, 3, 3}
	esperado = []int{3, 3, 3, 3, 3}
	verificarHeapSort(t, entrada, esperado, func(a, b int) int { return a - b })
}

func TestColaPrioridadHeapSortStrings(t *testing.T) {
	entrada := []string{"delta", "alpha", "charlie", "bravo"}
	esperado := []string{"alpha", "bravo", "charlie", "delta"}
	verificarHeapSort(t, entrada, esperado, strings.Compare)

	entrada = []string{"zeta", "epsilon", "beta", "gamma", "alpha"}
	esperado = []string{"alpha", "beta", "epsilon", "gamma", "zeta"}
	verificarHeapSort(t, entrada, esperado, strings.Compare)

	entrada = []string{}
	esperado = []string{}
	verificarHeapSort(t, entrada, esperado, strings.Compare)

	entrada = []string{"singleton"}
	esperado = []string{"singleton"}
	verificarHeapSort(t, entrada, esperado, strings.Compare)

	entrada = []string{"same", "same", "same"}
	esperado = []string{"same", "same", "same"}
	verificarHeapSort(t, entrada, esperado, strings.Compare)
}
