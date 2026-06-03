package main

import (
	app "algogram/main/app"
	comando "algogram/main/comando"
	utils "algogram/main/utils"
	"bufio"
	"os"
)

func main() {
	aplicacion := app.CrearApp()
	aplicacion.CargarUsuarios(utils.ObtenerArchivoSTDIn())
	manejador := comando.NuevoManejador(aplicacion)

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for scanner.Scan() {
		cmd := comando.ExtraerComando(scanner.Text())
		manejador.Ejecutar(cmd)
	}
}
