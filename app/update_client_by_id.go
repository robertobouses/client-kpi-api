package app

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateClientRequest struct {
	Name     *string
	LastName *string
	Email    *string
	Age      *int
	Birthday *string
}

func (a AppService) UpdateClientById(ctx *gin.Context, id uuid.UUID, req UpdateClientRequest) (string, error) {
	now := time.Now()
	var message string

	if req.Name == nil && req.LastName == nil && req.Email == nil && req.Age == nil && req.Birthday == nil {
		message = "Sin datos para actualizar"
		return message, errors.New("no hay datos para actualizar")
	}

	if req.Birthday != nil && req.Age != nil {
		message = "Se ignoró la edad, ya que la fecha de nacimiento fue proporcionada."
		birthday, err := time.Parse(time.RFC3339, *req.Birthday)
		if err != nil {
			return message, errors.New("formato de fecha de nacimiento inválido, use 'AAAA-MM-DDTHH:MM:SSZ'")
		}

		age := now.Year() - birthday.Year()
		if now.Month() < birthday.Month() || (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
			age--
		}
		req.Age = &age

		formattedBirthday := birthday.Format(time.RFC3339)
		req.Birthday = &formattedBirthday
	}

	if req.Age != nil && req.Birthday == nil {
		updatedYear := now.Year() - *req.Age
		birthday := time.Date(updatedYear, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		birthdayStr := birthday.Format(time.RFC3339)
		req.Birthday = &birthdayStr
	}

	if req.Birthday != nil && req.Age == nil {
		birthday, err := time.Parse(time.RFC3339, *req.Birthday)
		if err != nil {
			return "", errors.New("formato de fecha de nacimiento inválido, use 'AAAA-MM-DDTHH:MM:SSZ'")
		}

		age := now.Year() - birthday.Year()
		if now.Month() < birthday.Month() || (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
			age--
		}
		req.Age = &age
	}

	err := a.clientRepo.UpdateClientById(ctx, id, req)
	if err != nil {
		return "", err
	}

	return message, nil
}
