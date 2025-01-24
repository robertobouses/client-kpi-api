package app_test

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/mock"
)

type MockClientRepo struct {
	mock.Mock
}

func (m *MockClientRepo) InsertClients(req app.Client) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockClientRepo) ClientExistsById(id uuid.UUID) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockClientRepo) DeleteClientsById(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockClientRepo) QueryAllClients() ([]app.Client, error) {
	panic("not implemented")
}
func (m *MockClientRepo) QueryClientById(id uuid.UUID) (app.Client, error) {
	panic("not implemented")
}
func (m *MockClientRepo) UpdateClientById(ctx *gin.Context, id uuid.UUID, req app.UpdateClientRequest) error {
	args := m.Called(ctx, id, req)
	return args.Error(0)
}
