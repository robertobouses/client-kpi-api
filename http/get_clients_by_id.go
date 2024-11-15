package http

import (
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) GetClientById(ctx *gin.Context) {
	var request struct {
		Id uuid.UUID `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error en la solicitud. Asegúrate de enviar un JSON válido."})
		return
	}

	client, err := h.app.ListClientById(request.Id)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, client)

}
