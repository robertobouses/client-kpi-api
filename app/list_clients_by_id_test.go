package app_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListClientById(t *testing.T) {
	mockRepo := new(MockClientRepo)
	service := app.NewApp(mockRepo)

	mockClient := app.Client{
		Id:              uuid.New(),
		Name:            "Ramón",
		LastName:        "Vegas",
		Email:           "ra.vg@hotmali.com",
		Age:             30,
		Birthday:        time.Now().AddDate(-30, 0, 0),
		TelephoneNumber: "123456789",
	}

	tests := []struct {
		name           string
		inputID        uuid.UUID
		mockSetup      func()
		expectedResult app.Client
		expectedError  error
	}{
		{
			name:    "Cliente encontrado",
			inputID: mockClient.Id,
			mockSetup: func() {
				mockRepo.On("QueryClientById", mockClient.Id).Return(mockClient, nil).Once()
			},
			expectedResult: mockClient,
			expectedError:  nil,
		},
		{
			name:    "Cliente no encontrado",
			inputID: uuid.New(),
			mockSetup: func() {
				mockRepo.On("QueryClientById", mock.Anything).Return(app.Client{}, app.ErrClientNotFound).Once()
			},
			expectedResult: app.Client{},
			expectedError:  app.ErrClientNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			result, err := service.ListClientById(tt.inputID)

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
