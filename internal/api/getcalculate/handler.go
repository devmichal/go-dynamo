package getcalculate

import (
	"github.com/labstack/echo/v4"
	"musement.com/commission-service/internal/dto/response"
	"musement.com/commission-service/internal/repository"
	"net/http"
)

type Handler struct {
	repository repository.CommissionRepository
}

// Handle contains the handler logic
// @Summary Create a group of commission
// @Router /v1/calculate/{id} [get]
// @Accept json
// @Produce json
// @Param id path string true "The channel UUID"
// @Success 200 {object} response.Product
// @Failure 400 {object} api.ErrorAPI "Syntax error in provided payload"
// @Failure 401 {string} json "Unauthorized"
// @Failure 500 {string} json "Internal Server Error"
func New(repository repository.CommissionRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Handle(c echo.Context) error {

	id := c.Param("id")
	commission := h.repository.GetById(id)

	dto := response.Product{
		Id:        commission.Id,
		Name:      commission.Name,
		Price:     commission.Price,
		CreatedAt: commission.CreatedAt,
		Status:    commission.Status,
	}

	return c.JSON(http.StatusOK, dto)
}
