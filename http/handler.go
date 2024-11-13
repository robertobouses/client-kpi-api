package http

import (
	"github.com/robertobouses/client-kpi-api/app"
)

type App interface {
	ListAllClients() ([]app.Client, error)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
