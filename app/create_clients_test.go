package app_test

import (
	"testing"
	"time"

	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClients(t *testing.T) {
	tests := []struct {
		name          string
		client        app.Client
		mockSetup     func(mockRepo *MockClientRepo)
		expectedError string
		expectInsert  bool
	}{
		{
			name: "Edad y fecha de nacimiento coherentes",
			client: app.Client{
				Name:     "Juan Tijeras",
				Birthday: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
				Age:      34,
			},
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("InsertClients", mock.Anything).Return(nil).Once()
			},
			expectedError: "",
			expectInsert:  true,
		},
		{
			name: "Edad y fecha de nacimiento incoherentes",
			client: app.Client{
				Name:     "Ana Piedras",
				Birthday: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
				Age:      102,
			},
			mockSetup: func(mockRepo *MockClientRepo) {

			},
			expectedError: "la edad proporcionada no es coherente con la fecha de nacimiento",
			expectInsert:  false,
		},
		{
			name: "Edad no proporcionada (se calcula)",
			client: app.Client{
				Name:     "Carlos Papeles",
				Birthday: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
				Age:      0,
			},
			mockSetup: func(mockRepo *MockClientRepo) {
				mockRepo.On("InsertClients", mock.Anything).Return(nil).Once()
			},
			expectedError: "",
			expectInsert:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockClientRepo)
			tt.mockSetup(mockRepo)

			appService := app.NewApp(mockRepo)

			err := appService.CreateClients(tt.client)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.expectInsert {
				mockRepo.AssertCalled(t, "InsertClients", mock.Anything)
			} else {
				mockRepo.AssertNotCalled(t, "InsertClients", mock.Anything)
			}
		})
	}
}
