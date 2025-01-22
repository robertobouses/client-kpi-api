package app

import (
	"errors"
	"log"
	"time"
)

func (a AppService) CreateClients(req Client) error {

	now := time.Now().UTC()

	age := now.Year() - req.Birthday.Year()
	if now.Month() < req.Birthday.Month() || (now.Month() == req.Birthday.Month() && now.Day() < req.Birthday.Day()) {
		age--
	}

	if req.Age > 0 && req.Age != age {
		return errors.New("la edad proporcionada no es coherente con la fecha de nacimiento")
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
