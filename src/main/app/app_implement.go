package app

import (
	C "algogram/main/constantes"
	TDAPost "algogram/main/post"
	TDAUsuario "algogram/main/usuario"
	"bufio"
	"fmt"
	"os"
	TDADiccionario "tdas/diccionario"
)

type app struct {
	usuarioLogueado     TDAUsuario.Usuario
	usuariosRegistrados TDADiccionario.Diccionario[string, TDAUsuario.Usuario]
	postsPublicados     TDADiccionario.Diccionario[int, TDAPost.Post]
	idGenerator         TDAPost.IDGenerator
}

func CrearApp(generator TDAPost.IDGenerator) App {
	usuarios := TDADiccionario.CrearHash[string, TDAUsuario.Usuario](func(a, b string) bool { return a == b })
	postsPublicados := TDADiccionario.CrearHash[int, TDAPost.Post](func(a, b int) bool { return a == b })

	return &app{
		usuariosRegistrados: usuarios,
		usuarioLogueado:     nil,
		postsPublicados:     postsPublicados,
		idGenerator:         generator,
	}
}

// usuarioLogueadoONil valida que haya un usuario logueado.
// Centraliza el null-check para reducir repetición (DRY).
func (a *app) usuarioLogueadoONil() error {
	if a.usuarioLogueado == nil {
		return fmt.Errorf(C.ERROR_USUARIO_LOGOUT)
	}
	return nil
}

func (a *app) Login(nombreUsuario string) {
	if a.usuarioLogueado != nil {
		fmt.Println(C.ERROR_USUARIO_LOGIN)
		return
	}

	if a.usuariosRegistrados.Pertenece(nombreUsuario) {
		a.usuarioLogueado = a.usuariosRegistrados.Obtener(nombreUsuario)
		fmt.Println(C.BIENVENIDA, nombreUsuario)
	} else {
		fmt.Println(C.ERROR_USUARIO_NO_EXISTE)
	}
}

func (a *app) Logout() {
	if err := a.usuarioLogueadoONil(); err != nil {
		fmt.Println(err)
		return
	}

	a.usuarioLogueado = nil
	fmt.Println(C.DESPEDIDA)
}

func (a *app) ObtenerUsuarioLogueado() TDAUsuario.Usuario {
	return a.usuarioLogueado
}

func (a *app) CargarUsuarios(archivo string) error {
	archivoAbierto, err := os.Open(archivo)

	if err != nil {
		return fmt.Errorf(C.ERROR_ARCHIVO)
	}

	defer archivoAbierto.Close()
	scanner := bufio.NewScanner(archivoAbierto)

	id := 0
	for scanner.Scan() {
		nombre := scanner.Text()
		if nombre == "" {
			continue
		}
		nuevoUsuario := TDAUsuario.CrearUsuario(nombre, id)
		a.usuariosRegistrados.Guardar(nombre, nuevoUsuario)
		id++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(C.ERROR_ARCHIVO)
	}

	return nil
}

func (a *app) ObtenerUsuariosRegistrados() TDADiccionario.Diccionario[string, TDAUsuario.Usuario] {
	return a.usuariosRegistrados
}

func (a *app) Publicar(contenido string) {
	if err := a.usuarioLogueadoONil(); err != nil {
		fmt.Println(err)
		return
	}

	autor := a.usuarioLogueado.ObtenerNombre()
	autor_id := a.usuarioLogueado.ObtenerId()
	nuevoPost := TDAPost.CrearPost(a.idGenerator, autor, contenido, autor_id)
	iteradorUsuarios := a.usuariosRegistrados.Iterador()

	for iteradorUsuarios.HaySiguiente() {
		usuarioActual, _ := iteradorUsuarios.VerActual()

		if autor != usuarioActual {
			feedUsuarioActual := a.usuariosRegistrados.Obtener(usuarioActual).ObtenerFeed()
			feedUsuarioActual.Encolar(nuevoPost)
		}

		iteradorUsuarios.Siguiente()
	}

	a.postsPublicados.Guardar(nuevoPost.ObtenerID(), nuevoPost)
	fmt.Println(C.POST_PUBLICADO)
}

func (a *app) VerSiguienteFeed() {
	if err := a.usuarioLogueadoONil(); err != nil {
		fmt.Println(err)
		return
	}

	feed := a.usuarioLogueado.ObtenerFeed()

	if feed.EstaVacia() {
		fmt.Println(C.ERROR_NO_HAY_POSTS)
		return
	}

	post := feed.Desencolar()
	fmt.Println("Post ID", post.ObtenerID())
	fmt.Println(post.ObtenerAutor(), "dijo:", post.ObtenerContenido())
	fmt.Println("Likes:", post.CantidadLikes())
}

func (a *app) LikearPost(id int) {
	if err := a.usuarioLogueadoONil(); err != nil {
		fmt.Println(err)
		return
	}

	if !a.postsPublicados.Pertenece(id) {
		fmt.Println(C.ERROR_LIKEAR_POST)
		return
	}

	post := a.postsPublicados.Obtener(id)
	nombreUsuario := a.usuarioLogueado.ObtenerNombre()

	if !post.UsuarioYaLikeo(nombreUsuario) {
		post.Likear(nombreUsuario)
	}

	fmt.Println(C.POST_LIKEADO)
}

func (a *app) MostrarLikes(id int) {
	if !a.postsPublicados.Pertenece(id) {
		fmt.Println(C.ERROR_MOSTRAR_LIKES)
		return
	}

	post := a.postsPublicados.Obtener(id)

	if post.CantidadLikes() == 0 {
		fmt.Println(C.ERROR_MOSTRAR_LIKES)
		return
	}

	fmt.Println("El post tiene", post.CantidadLikes(), "likes:")
	usuariosLikes := post.ObtenerUsuariosLikes()

	for i := 0; i < len(usuariosLikes); i++ {
		fmt.Printf("\t%s\n", usuariosLikes[i])
	}
}
