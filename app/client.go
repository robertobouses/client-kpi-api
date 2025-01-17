package app

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id              uuid.UUID
	Name            string
	LastName        string
	Email           string
	Age             int
	Birthday        time.Time
	TelephoneNumber string
}
