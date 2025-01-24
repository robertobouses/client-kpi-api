package app_test

import (
	"errors"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateClientById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockClientRepo)
	service := app.NewApp(mockRepo)
	ctx := &gin.Context{}

	id := uuid.New()
	now := time.Now()
	birthdayStr := now.AddDate(-30, 0, 0).Format(time.RFC3339)

	tests := []struct {
		name          string
		inputRequest  app.UpdateClientRequest
		mockSetup     func()
		expectedMsg   string
		expectedError string
	}{
		{
			name: "Sin datos para actualizar",
			inputRequest: app.UpdateClientRequest{
				Name: nil, LastName: nil, Email: nil, Age: nil, Birthday: nil,
			},
			mockSetup:     func() {},
			expectedMsg:   "Sin datos para actualizar",
			expectedError: "no hay datos para actualizar",
		},
		{
			name: "Proporcionados Edad(ignorar) y Fecha de nacimiento",
			inputRequest: app.UpdateClientRequest{
				Birthday: &birthdayStr,
				Age:      new(int),
			},
			mockSetup: func() {
				mockRepo.On("UpdateClientById", mock.Anything, id, mock.Anything).Return(nil).Once()
			},
			expectedMsg:   "Se ignoró la edad, ya que la fecha de nacimiento fue proporcionada.",
			expectedError: "",
		},
		{
			name: "Formato de Fecha de nacimeinto inválido",
			inputRequest: app.UpdateClientRequest{
				Birthday: new(string),
			},
			mockSetup:     func() {},
			expectedMsg:   "",
			expectedError: "formato de fecha de nacimiento inválido, use 'AAAA-MM-DDTHH:MM:SSZ'",
		},
		{
			name: "Proporcionada solo Edad, calcular Fecha nacimiento",
			inputRequest: app.UpdateClientRequest{
				Age: new(int),
			},
			mockSetup: func() {
				mockRepo.On("UpdateClientById", mock.Anything, id, mock.Anything).Return(nil).Once()
			},
			expectedMsg:   "",
			expectedError: "",
		},
		{
			name: "Error en el Repo",
			inputRequest: app.UpdateClientRequest{
				Name: new(string),
			},
			mockSetup: func() {
				mockRepo.On("UpdateClientById", mock.Anything, id, mock.Anything).Return(errors.New("repository error")).Once()
			},
			expectedMsg:   "",
			expectedError: "repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			msg, err := service.UpdateClientById(ctx, id, tt.inputRequest)

			assert.Equal(t, tt.expectedMsg, msg)
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
