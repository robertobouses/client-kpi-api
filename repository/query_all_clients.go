package repository

import (
	"log"

	"github.com/robertobouses/client-kpi-api/app"
)

func (r *repository) QueryAllClients() ([]app.Client, error) {

	rows, err := r.queryAllClients.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []app.Client
	for rows.Next() {
		var client app.Client
		if err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.LastName,
			&client.Email,
			&client.Age,
			&client.Birthday,
			&client.TelephoneNumber,
		); err != nil {
			log.Printf("Error al escanear las filas QueryAllClients: %v", err)
			return nil, err
		}
		clients = append(clients, client)

	}

	return clients, nil
}
