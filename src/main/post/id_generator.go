package post

// IDGenerator genera identificadores únicos para posts
type IDGenerator interface {
	Next() int
}

// idGenerator es la implementación de IDGenerator
type idGenerator struct {
	contador int
}

// NewIDGenerator crea un nuevo generador de IDs iniciando desde -1
// Cada llamada a Next() incrementa y retorna el ID siguiente
func NewIDGenerator() IDGenerator {
	return &idGenerator{contador: -1}
}

// Next retorna el siguiente ID disponible
func (g *idGenerator) Next() int {
	g.contador++
	return g.contador
}
