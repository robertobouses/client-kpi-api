package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

type UpdateClientRequest struct {
	Name            *string `json:"name"`
	LastName        *string `json:"last_name"`
	Email           *string `json:"email"`
	Age             *int    `json:"age"`
	Birthday        *string `json:"birthday"`
	TelephoneNumber *string `json:"telephone_number"`
}

func (h Handler) PutClientsById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en los datos de entrada: " + err.Error()})
		return
	}
	appReq := app.UpdateClientRequest{
		Name:            req.Name,
		LastName:        req.LastName,
		Email:           req.Email,
		Age:             req.Age,
		Birthday:        req.Birthday,
		TelephoneNumber: req.TelephoneNumber,
	}
	message, err := h.app.UpdateClientById(c, id, appReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el cliente: " + err.Error()})
		return
	}

	if message != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"cliente": "Cliente actualizado exitosamente",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Cliente actualizado exitosamente"})
	}
}
