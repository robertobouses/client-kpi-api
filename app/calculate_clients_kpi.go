package app

import "math"

type KPI struct {
	AverageAge float64 `json:"average_age"`
	StdDevAge  float64 `json:"std_dev_age"`
}

func (a AppService) CalculateClientsKPI(clients []Client) (KPI, error) {
	var averageAge, stdDevAge, sum, sumSquared float64
	n := float64(len(clients))

	if n == 0 {
		return KPI{}, nil
	}

	for _, client := range clients {
		sum += float64(client.Age)
	}

	averageAge = sum / n

	for _, client := range clients {
		sumSquared += math.Pow(float64(client.Age)-averageAge, 2)
	}

	stdDevAge = math.Sqrt(sumSquared / n)

	return KPI{
		AverageAge: averageAge,
		StdDevAge:  stdDevAge,
	}, nil
}
