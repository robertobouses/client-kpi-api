package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateClientRequest struct {
	Name            *string
	LastName        *string
	Email           *string
	Age             *int
	Birthday        *string
	TelephoneNumber *string
}

func (a AppService) UpdateClientById(ctx *gin.Context, id uuid.UUID, req UpdateClientRequest) (string, error) {

	var message string

	if req.Name == nil && req.LastName == nil && req.Email == nil && req.Age == nil && req.Birthday == nil {
		message = "Sin datos para actualizar"
		return message, ErrNoDataToUpdate
	}

	if req.Birthday != nil && req.Age != nil {
		message = "Se ignoró la edad, ya que la fecha de nacimiento fue proporcionada."
		birthday, err := time.Parse(time.RFC3339, *req.Birthday)
		if err != nil {
			return message, ErrInvalidBirthdayFormat
		}

		calculatedAge := CalculateAge(birthday)
		req.Age = &calculatedAge

		formattedBirthday := birthday.Format(time.RFC3339)
		req.Birthday = &formattedBirthday
	}

	if req.Age != nil && req.Birthday == nil {
		now := time.Now()
		updatedYear := now.Year() - *req.Age
		birthday := time.Date(updatedYear, now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		birthdayStr := birthday.Format(time.RFC3339)
		req.Birthday = &birthdayStr
	}

	if req.Birthday != nil && req.Age == nil {
		birthday, err := time.Parse(time.RFC3339, *req.Birthday)
		if err != nil {
			return "", ErrInvalidBirthdayFormat
		}

		calculatedAge := CalculateAge(birthday)
		req.Age = &calculatedAge
	}

	err := a.clientRepo.UpdateClientById(ctx, id, req)
	if err != nil {
		return "", err
	}

	return message, nil
}
