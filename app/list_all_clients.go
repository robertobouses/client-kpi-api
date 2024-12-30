package app

import "log"

func (a AppService) ListAllClients() ([]Client, error) {
	clients, err := a.clientRepo.QueryAllClients()
	if err != nil {
		log.Println("Error al extraer ListAllClients", err)
		return []Client{}, err
	}
	return clients, nil
}
