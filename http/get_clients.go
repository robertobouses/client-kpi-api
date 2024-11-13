package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetClients(ctx *gin.Context) {
	clients, err := h.app.ListAllClients()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, clients)

}
