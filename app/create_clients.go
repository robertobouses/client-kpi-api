package app

import (
	"log"
	"time"
)

func (a AppService) CreateClients(req Client) error {

	now := time.Now()
	age := now.Year() - req.Birthday.Year()
	if now.Month() < req.Birthday.Month() || (now.Month() == req.Birthday.Month() && now.Day() < req.Birthday.Day()) {
		age--
	}

	req.Age = age

	log.Printf("Valores del cliente creado: %+v\n", req)

	err := a.clientRepo.InsertClients(req)
	if err != nil {
		log.Println("Error en InsertClients:", err)
		return err
	}
	return nil
}
