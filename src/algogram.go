package main

import (
	"bufio"
	"fmt"
	"os"
	app "algogram/main/app"
	comando "algogram/main/comando"
	post "algogram/main/post"
	utils "algogram/main/utils"
)

func main() {
	// Crear generador de IDs (inyección de dependencia)
	idGenerator := post.NewIDGenerator()

	// Crear aplicación con el generador inyectado
	aplicacion := app.CrearApp(idGenerator)

	// Cargar usuarios desde archivo
	archivoPath := utils.ObtenerArchivoSTDIn()
	if err := aplicacion.CargarUsuarios(archivoPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Crear manejador de comandos
	manejador := comando.NuevoManejador(aplicacion)

	// Procesar comandos desde entrada estándar
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for scanner.Scan() {
		cmd := comando.ExtraerComando(scanner.Text())
		if !manejador.Ejecutar(cmd) {
			break
		}
	}
}
