package app

import "errors"

var (
	ErrClientNotFound        = errors.New("cliente no existe")
	ErrClientDeletion        = errors.New("error al eliminar cliente")
	ErrListClients           = errors.New("error al listar clientes, error db")
	ErrInconsistentAge       = errors.New("la edad proporcionada no es coherente con la fecha de nacimiento")
	ErrRepoUpdateClient      = errors.New("error al actualizar el cliente en el repositorio")
	ErrNoDataToUpdate        = errors.New("no hay datos para actualizar")
	ErrInvalidBirthdayFormat = errors.New("formato de fecha de nacimiento inválido, use 'AAAA-MM-DDTHH:MM:SSZ'")
)
