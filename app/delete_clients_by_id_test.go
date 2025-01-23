package app_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteClientsById(t *testing.T) {
	tests := []struct {
		name           string
		clientID       uuid.UUID
		mockSetup      func(mockRepo *MockClientRepo)
		expectedError  string
		expectDeletion bool
	}{
		{
			name:     "Cliente existente, se elimina correctamente",
			clientID: uuid.New(),
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("ClientExistsById", mock.Anything).Return(true, nil).Once()
				mockRepo.On("DeleteClientsById", mock.Anything).Return(nil).Once()
			},
			expectedError:  "",
			expectDeletion: true,
		},
		{
			name:     "Cliente no existente",
			clientID: uuid.New(),
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("ClientExistsById", mock.Anything).Return(false, nil).Once()
			},
			expectedError:  "el cliente no existe",
			expectDeletion: false,
		},
		{
			name:     "Error al comprobar existencia del cliente",
			clientID: uuid.New(),
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("ClientExistsById", mock.Anything).Return(false, errors.New("error en la base de datos")).Once()
			},
			expectedError:  "error en la base de datos",
			expectDeletion: false,
		},
		{
			name:     "Error al eliminar el cliente",
			clientID: uuid.New(),
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("ClientExistsById", mock.Anything).Return(true, nil).Once()
				mockRepo.On("DeleteClientsById", mock.Anything).Return(errors.New("error al eliminar")).Once()
			},
			expectedError:  "error al eliminar",
			expectDeletion: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockClientRepo)
			tt.mockSetup(mockRepo)

			appService := app.NewApp(mockRepo)

			err := appService.DeleteClientsById(tt.clientID)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.expectDeletion {
				mockRepo.AssertCalled(t, "DeleteClientsById", tt.clientID)
			} else {
				mockRepo.AssertNotCalled(t, "DeleteClientsById", tt.clientID)
			}
		})
	}
}
