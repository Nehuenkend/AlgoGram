package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	C "tp2/constantes"
)

func Modulo(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func StrCmp(a, b string) int {
	if b < a {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func ObtenerArchivoSTDIn() string {
	parametros := os.Args

	if len(parametros) < 2 {
		fmt.Println(C.ERROR_ARCHIVO_STDIN)
		os.Exit(1)
	}

	return parametros[1]
}

func ObtenerIDPostDesdeComando(parametro, mensajeError string) (int, error) {
	idPost, err := strconv.Atoi(parametro)

	if err != nil {
		return -1, errors.New(mensajeError)
	}

	return idPost, nil
}
