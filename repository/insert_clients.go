package repository

import (
	"log"

	"github.com/robertobouses/client-kpi-api/app"
)

func (r *repository) InsertClients(req app.Client) error {
	_, err := r.insertClients.Exec(
		req.Name,
		req.LastName,
		req.Email,
		req.Age,
		req.Birthday,
		req.TelephoneNumber,
	)
	if err != nil {
		log.Print("Error en InsertClients repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
