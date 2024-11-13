package app

type ClientRepository interface {
	QueryAllClients() ([]Client, error)
}

func NewApp(clientRepository ClientRepository) AppService {
	return AppService{
		clientRepo: clientRepository,
	}
}

type AppService struct {
	clientRepo ClientRepository
}
