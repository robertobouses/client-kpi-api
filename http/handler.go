package http

import (
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

type App interface {
	ListAllClients() ([]app.Client, error)
	ListClientById(id uuid.UUID) (app.Client, error)
	CreateClients(req app.Client) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
