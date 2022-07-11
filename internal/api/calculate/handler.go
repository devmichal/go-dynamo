package calculate

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"musement.com/commission-service/internal/api"
	"musement.com/commission-service/internal/dto/request"
	"musement.com/commission-service/internal/model"
	"musement.com/commission-service/internal/repository"
	"net/http"
	"time"
)

type Handler struct {
	repository repository.CommissionRepository
}

func New(repository repository.CommissionRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) Handler(c echo.Context) error {
	request := new(request.Product)

	_ = (&echo.DefaultBinder{}).Bind(request, c)

	if errParam := c.Validate(request); errParam != nil {

		return c.JSON(
			http.StatusBadRequest,
			&api.ErrorAPI{Message: fmt.Sprintf("Request Path %s", errParam.Error())},
		)
	}

	id := uuid.New()

	err := h.repository.CreateCommission(&model.Commission{
		Id:        id.String(),
		Price:     request.Price,
		Status:    request.Status,
		Name:      request.Name,
		CreatedAt: time.Now().Format("2006-01-02 15:02"),
	})

	if err != nil {
		panic(err)
	}
	println(request.Name)

	return c.JSON(http.StatusOK, id.String())
}
