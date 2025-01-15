package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

func (r *repository) QueryClientById(id uuid.UUID) (app.Client, error) {

	row := r.queryClientById.QueryRow(id)
	var client app.Client
	if err := row.Scan(
		&client.Id,
		&client.Name,
		&client.LastName,
		&client.Email,
		&client.Age,
		&client.Birthday,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontró el cliente con clientid: %v", id)
			return app.Client{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return app.Client{}, err
	}

	return client, nil
}
