package main

import (
	"bufio"
	"fmt"
	"os"
	C "tp2/constantes"
	TDAApp "tp2/tdas/app"
	TDAComando "tp2/tdas/comando"
	"tp2/utils"
)

func main() {
	app := TDAApp.CrearApp()
	archivo := utils.ObtenerArchivoSTDIn()
	app.CargarUsuarios(archivo)

	var comandos = map[string]func(param string){
		C.COMANDO_LOGIN: func(param string) {
			app.Login(param)
		},
		C.COMANDO_LOGOUT: func(param string) {
			app.Logout()
		},
		C.COMANDO_PUBLICAR: func(param string) {
			app.Publicar(param)
		},
		C.COMANDO_VER_SIGUIENTE_FEED: func(param string) {
			app.VerSiguienteFeed()
		},
		C.COMANDO_LIKEAR_POST: func(param string) {
			idPost, err := utils.ObtenerIDPostDesdeComando(param, C.ERROR_LIKEAR_POST)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			app.LikearPost(idPost)
		},
		C.COMANDO_MOSTRAR_LIKES: func(param string) {
			idPost, err := utils.ObtenerIDPostDesdeComando(param, C.ERROR_MOSTRAR_LIKES)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			app.MostrarLikes(idPost)
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		comando := TDAComando.ExtraerComando(scanner.Text())

		cmdFunc, ok := comandos[comando.Tipo()]
		if ok {
			cmdFunc(comando.Parametro())
		} else {
			fmt.Println(C.COMANDO_NO_RECONOCIDO)
		}
	}
}
