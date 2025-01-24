package http

import (
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

func (h Handler) GetClientsById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": app.ErrInvalidID})
		return
	}

	client, err := h.app.ListClientById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": app.ErrClientNotFound})
		return
	}

	ctx.JSON(nethttp.StatusOK, client)

}
