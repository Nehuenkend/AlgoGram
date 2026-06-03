package comando

import (
	app "algogram/main/app"
	constantes "algogram/main/constantes"
	utils "algogram/main/utils"
	"fmt"
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

// ejecutarConID es un helper genérico que parsea un ID y ejecuta una función.
func (m *ManejadorComandos) ejecutarConID(param string, errorMsg string, fn func(int)) {
	id, err := utils.ObtenerIDPostDesdeComando(param, errorMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fn(id)
}

// likearPost parsea el ID del post y delega a app.LikearPost.
func (m *ManejadorComandos) likearPost(param string) {
	m.ejecutarConID(param, constantes.ERROR_LIKEAR_POST, m.app.LikearPost)
}

// mostrarLikes parsea el ID del post y delega a app.MostrarLikes.
func (m *ManejadorComandos) mostrarLikes(param string) {
	m.ejecutarConID(param, constantes.ERROR_MOSTRAR_LIKES, m.app.MostrarLikes)
}
