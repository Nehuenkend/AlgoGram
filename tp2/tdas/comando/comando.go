package tdas

type Comando interface {
	// Tipo devuelve el nombre del comando.
	Tipo() string

	// Parametros devuelve un parametro del comando asociado.
	Parametro() string
}
