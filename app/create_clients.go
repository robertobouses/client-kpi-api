package app

import (
	"log"
	"time"
)

func (a AppService) CreateClients(req Client) error {

	age := CalculateAge(req.Birthday)

	if req.Age > 0 && req.Age != age {
		return ErrInconsistentAge
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

func CalculateAge(birthday time.Time) int {
	now := time.Now().UTC()

	age := now.Year() - birthday.Year()
	if now.Month() < birthday.Month() || (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
		age--
	}
	return age
}
