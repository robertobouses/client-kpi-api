package http

import (
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/client-kpi-api/app"
)

func (h Handler) GetClientsKPI(ctx *gin.Context) {

	clients, err := h.app.ListAllClients()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(clients) == 0 {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"average_age": 0,
			"std_dev_age": 0,
		})
		return
	}

	kpi, err := h.app.CalculateClientsKPI(clients)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": app.ErrKpiCalculationFailed.Error()})
		return
	}
	ctx.JSON(nethttp.StatusOK, gin.H{
		"average_age": kpi.AverageAge,
		"std_dev_age": kpi.StdDevAge,
	})
}
