package calculateQuery

import (
	"github.com/labstack/echo/v4"
	"musement.com/commission-service/internal/repository"
	"net/http"
)

type Handler struct {
	repository repository.CommissionRepository
}

func New(repository repository.CommissionRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Handle(c echo.Context) error {

	id := c.Param("id")
	commission := h.repository.GetById(id)

	return c.JSON(http.StatusOK, commission)
}
