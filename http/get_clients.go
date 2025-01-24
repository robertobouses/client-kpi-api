package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/client-kpi-api/app"
)

func (h Handler) GetClients(ctx *gin.Context) {
	clients, err := h.app.ListAllClients()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": app.ErrListClients})
		return
	}

	ctx.JSON(nethttp.StatusOK, clients)

}
