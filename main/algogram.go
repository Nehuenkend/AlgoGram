package main

import (
	"bufio"
	app "main/src/app"
	comando "main/src/comando"
	utils "main/src/utils"
	"os"
)

func main() {
	aplicacion := app.CrearApp()
	aplicacion.CargarUsuarios(utils.ObtenerArchivoSTDIn())
	manejador := comando.NuevoManejador(aplicacion)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := comando.ExtraerComando(scanner.Text())
		manejador.Ejecutar(cmd)
	}
}
