package app

import "math"

func (a AppService) CalculateClientsKPI(clients []Client) (float64, float64) {
	var averageAge, stdDevAge, sum, sumSquared float64
	n := float64(len(clients))

	for _, client := range clients {
		sum += float64(client.Age)
	}

	averageAge = sum / n

	for _, client := range clients {
		sumSquared += math.Pow(float64(client.Age)-averageAge, 2)
	}

	stdDevAge = math.Sqrt(sumSquared / n)

	return averageAge, stdDevAge
}
