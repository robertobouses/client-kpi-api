package http

import (
	"log"
	"net/http"
	nethttp "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/client-kpi-api/app"
)

type PostClientsRequest struct {
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

func (h Handler) PostClients(c *gin.Context) {
	var req PostClientsRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	client := app.Client{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Birthday: req.Birthday,
	}
	err := h.app.CreateClients(client)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar la app desde http": err.Error()})
		return
	}
	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "cliente insertado correctamente"})
}
