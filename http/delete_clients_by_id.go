package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

func (h Handler) DeleteClientsById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": app.ErrInvalidID})
		return
	}

	err = h.app.DeleteClientsById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": app.ErrClientNotFound})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado exitosamente"})
}
