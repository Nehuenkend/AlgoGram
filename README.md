# TP2 Algoritmos y Estructuras de Datos FiUBA - Algogram 📱

Simulación de una red social minimalista implementada en **Go**, desarrollada como trabajo práctico universitario. Algogram permite a usuarios registrados publicar posts, ver un feed personalizado, dar likes y consultar quiénes likearon cada publicación.

---

## ¿Qué hace?

El sistema simula el backend de una red social con las siguientes funcionalidades:

- **Login / Logout** de usuarios registrados
- **Publicar posts** con contenido de texto
- **Feed personalizado**: cada usuario ve los posts de los demás, ordenados por *afinidad* (qué tan cercano es el ID del autor al propio) y luego por ID de post
- **Likear posts** (sin duplicados por usuario)
- **Ver likes** de un post, con la lista de usuarios que likearon ordenada alfabéticamente

El feed está implementado con una **cola de prioridad (max-heap)** que prioriza posts de usuarios "afines" (aquellos cuyo ID numérico es más cercano al del usuario que consulta).

---

## Estructuras de datos utilizadas

| Estructura | Uso |
|---|---|
| Hash / Diccionario | Usuarios registrados, posts publicados, likes por post |
| Cola de prioridad (Heap) | Feed de cada usuario (ordenado por afinidad), lista de usuarios que likearon (orden alfabético) |

---

## Cómo correrlo

### Requisitos

- [Go](https://golang.org/dl/) 1.18 o superior

### Estructura esperada del proyecto

```
tp2/
├── algogram.go
├── tdas/
│   ├── app/
│   ├── post/
│   ├── usuario/
│   ├── comando/
│   ├── diccionario/
│   └── cola_prioridad/
├── constantes/
└── utils/
```

### Ejecución

El programa recibe por `stdin` el nombre del archivo de usuarios y luego lee comandos línea a línea:

```bash
go run algogram.go < entrada.txt
```

O bien pasando el archivo de usuarios como primer argumento (según la implementación de `utils.ObtenerArchivoSTDIn()`).

### Formato del archivo de usuarios

Un nombre de usuario por línea:

```
juan
maria
pedro
ana
```

### Comandos disponibles

| Comando | Descripción |
|---|---|
| `login <usuario>` | Inicia sesión con el usuario dado |
| `logout` | Cierra la sesión actual |
| `publicar <contenido>` | Publica un post con el texto dado |
| `ver_siguiente_feed` | Muestra el siguiente post del feed |
| `likear_post <id>` | Da like al post con el ID indicado |
| `mostrar_likes <id>` | Muestra cuántos y quiénes likearon el post |

### Ejemplo de ejecución

```
> login juan
Bienvenido juan
> publicar Hola mundo!
Post publicado
> logout
Adios
> login maria
Bienvenido maria
> ver_siguiente_feed
Post ID 0
juan dijo: Hola mundo!
Likes: 0
> likear_post 0
Post likeado
> mostrar_likes 0
El post tiene 1 likes:
	maria
> logout
Adios
```

---

## Decisiones de diseño

- El **feed** es un heap de mínimos por distancia de afinidad: cuanto menor la diferencia entre el ID del autor y el ID del usuario, mayor prioridad tiene el post. En caso de empate, se prioriza el post más reciente (mayor ID).
- Los **likes** se almacenan tanto en un diccionario (para chequear duplicados en O(1)) como en un heap (para devolverlos ordenados alfabéticamente).
- Los **usuarios** y **posts** se acceden mediante diccionarios hash, lo que garantiza búsqueda y guardado en O(1) amortizado.
