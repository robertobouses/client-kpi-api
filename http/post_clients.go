package http

import (
	"log"
	"net/http"
	nethttp "net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/client-kpi-api/app"
)

type PostClientsRequest struct {
	Name            string    `json:"name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Birthday        time.Time `json:"birthday"`
	TelephoneNumber string    `json:"telephone_number"`
}

func (h Handler) PostClients(c *gin.Context) {
	var req PostClientsRequest

	if err := c.BindJSON(&req); err != nil {
		log.Printf("error parsing request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de entrada inválido"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.LastName == "" || req.Email == "" || req.Birthday.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, LastName, Email y Fecha de nacimiento son obligatorios"})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de email inválido"})
		return
	}
	if req.Birthday.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha de nacimiento no puede ser futura"})
		return
	}

	if !isValidPhoneNumber(req.TelephoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Número de teléfono inválido"})
		return
	}
	client := app.Client{
		Name:            req.Name,
		LastName:        req.LastName,
		Email:           req.Email,
		Birthday:        req.Birthday,
		TelephoneNumber: req.TelephoneNumber,
	}
	err := h.app.CreateClients(client)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar la app desde http": err.Error()})
		return
	}
	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "cliente insertado correctamente"})
}

func isValidEmail(email string) bool {
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
func isValidPhoneNumber(phone string) bool {
	regex := `^\d{7,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(phone)
}
