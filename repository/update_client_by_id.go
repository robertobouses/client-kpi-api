package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/client-kpi-api/app"
)

func (r *repository) UpdateClientById(ctx *gin.Context, id uuid.UUID, req app.UpdateClientRequest) error {
	_, err := r.updateClientById.ExecContext(
		ctx,
		req.Name,
		req.LastName,
		req.Email,
		req.Age,
		req.Birthday,
		id,
	)
	if err != nil {
		ctx.Error(err)
		return err
	}

	return nil
}
