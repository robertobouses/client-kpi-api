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

	ErrInvalidRequestFormat     = errors.New("formato de entrada inválido")
	ErrMandatoryFieldsMissing   = errors.New("name, lastName, email y fecha de nacimiento son obligatorios")
	ErrInvalidEmailFormat       = errors.New("formato de email inválido")
	ErrFutureBirthdayNotAllowed = errors.New("la fecha de nacimiento no puede ser futura")
	ErrInvalidPhoneNumber       = errors.New("número de teléfono inválido")
	ErrAppCommunication         = errors.New("error al llamar la app desde http")
	ErrInvalidID                = errors.New("ID inválido")
	ErrClientUpdate             = errors.New("no se pudo actualizar el cliente")
	ErrKpiCalculationFailed     = errors.New("error al calcular KPI")
)
