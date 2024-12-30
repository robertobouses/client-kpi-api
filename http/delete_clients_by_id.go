package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) DeleteClientsById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID proporcionado no es válido."})
		return
	}

	err = h.app.DeleteClientsById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado exitosamente"})
}
