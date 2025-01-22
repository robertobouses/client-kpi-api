package app_test

import (
	"testing"

	"github.com/robertobouses/client-kpi-api/app"
	"github.com/stretchr/testify/assert"
)

func TestCalculateClientsKPI(t *testing.T) {
	tests := []struct {
		name     string
		clients  []app.Client
		expected app.KPI
	}{
		{
			name: "Clientes con edades variadas",
			clients: []app.Client{
				{Age: 25},
				{Age: 30},
				{Age: 35},
				{Age: 40},
			},
			expected: app.KPI{
				AverageAge: 32.5,
				StdDevAge:  5.5901699437494745,
			},
		},
		{
			name:    "Sin clientes",
			clients: []app.Client{},
			expected: app.KPI{
				AverageAge: 0,
				StdDevAge:  0,
			},
		},
		{
			name: "Un solo cliente",
			clients: []app.Client{
				{Age: 50},
			},
			expected: app.KPI{
				AverageAge: 50,
				StdDevAge:  0,
			},
		},
		{
			name: "Clientes con la misma edad",
			clients: []app.Client{
				{Age: 20},
				{Age: 20},
				{Age: 20},
			},
			expected: app.KPI{
				AverageAge: 20,
				StdDevAge:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockClientRepo)
			appService := app.NewApp(mockRepo)

			result, err := appService.CalculateClientsKPI(tt.clients)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected.AverageAge, result.AverageAge, 0.0001, "Average age mismatch")
			assert.InDelta(t, tt.expected.StdDevAge, result.StdDevAge, 0.0001, "Standard deviation mismatch")
		})
	}
}
