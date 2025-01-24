package app_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
)

func TestListAllClients(t *testing.T) {
	mockRepo := new(MockClientRepo)
	service := app.NewApp(mockRepo)

	mockClients := []app.Client{
		{
			Id:              uuid.New(),
			Name:            "Pilar",
			LastName:        "Pilares",
			Email:           "pi.pilares@gmail.com",
			Age:             30,
			Birthday:        time.Now().AddDate(-30, 0, 0),
			TelephoneNumber: "123456789",
		},
		{
			Id:              uuid.New(),
			Name:            "Carmelo",
			LastName:        "Carmen",
			Email:           "ca.carmen@example.com",
			Age:             25,
			Birthday:        time.Now().AddDate(-25, 0, 0),
			TelephoneNumber: "987654321",
		},
	}

	tests := []struct {
		name           string
		mockSetup      func()
		expectedResult []app.Client
		expectedError  error
	}{
		{
			name: "Listar clientes correctamente",
			mockSetup: func() {
				mockRepo.On("QueryAllClients").Return(mockClients, nil).Once()
			},
			expectedResult: mockClients,
			expectedError:  nil,
		},
		{
			name: "Error al listar clientes",
			mockSetup: func() {
				mockRepo.On("QueryAllClients").Return([]app.Client{}, errors.New("database error")).Once()
			},
			expectedResult: []app.Client{},
			expectedError:  errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			result, err := service.ListAllClients()

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
