package src

import (
	"fmt"
	app "main/src/app"
	constantes "main/src/constantes"
	utils "main/src/utils"
)

// ManejadorComandos encapsula la asociación entre comandos y su ejecución.
type ManejadorComandos struct {
	app      app.App
	comandos map[string]func(string)
}

// NuevoManejador crea un manejador con todos los comandos registrados.
func NuevoManejador(a app.App) *ManejadorComandos {
	m := &ManejadorComandos{app: a}
	m.comandos = map[string]func(string){
		constantes.COMANDO_LOGIN:              a.Login,
		constantes.COMANDO_LOGOUT:             func(string) { a.Logout() },
		constantes.COMANDO_PUBLICAR:           a.Publicar,
		constantes.COMANDO_VER_SIGUIENTE_FEED: func(string) { a.VerSiguienteFeed() },
		constantes.COMANDO_LIKEAR_POST:        m.likearPost,
		constantes.COMANDO_MOSTRAR_LIKES:      m.mostrarLikes,
	}
	return m
}

// Ejecutar invoca el comando correspondiente con su parámetro.
func (m *ManejadorComandos) Ejecutar(cmd Comando) {
	if fn, ok := m.comandos[cmd.Tipo()]; ok {
		fn(cmd.Parametro())
	} else {
		fmt.Println(constantes.COMANDO_NO_RECONOCIDO)
	}
}

// likearPost parsea el ID y llama al método de la app.
func (m *ManejadorComandos) likearPost(param string) {
	id, err := utils.ObtenerIDPostDesdeComando(param, constantes.ERROR_LIKEAR_POST)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.app.LikearPost(id)
}

// mostrarLikes parsea el ID y llama al método de la app.
func (m *ManejadorComandos) mostrarLikes(param string) {
	id, err := utils.ObtenerIDPostDesdeComando(param, constantes.ERROR_MOSTRAR_LIKES)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.app.MostrarLikes(id)
}
