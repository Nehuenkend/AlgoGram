package tdas

import "strings"

type comando struct {
	tipo      string
	parametro string
}

func ExtraerComando(linea string) Comando {
	var tipo string
	var parametros string

	partes := strings.Fields(linea)

	if len(partes) > 0 {
		tipo = partes[0]
		parametros = strings.Join(partes[1:], " ")
	}

	return comando{tipo: tipo, parametro: parametros}
}

func (c comando) Tipo() string {
	return c.tipo
}

func (c comando) Parametro() string {
	return c.parametro
}
