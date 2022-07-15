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

// Handle contains the handler logic
// @Summary Create a group of commission
// @Router /v1/calculate [post]
// @Accept json
// @Produce json
// @Param body body request.Product true "Array of uuid groups"
// @Success 204
// @Failure 400 {object} api.ErrorAPI "Syntax error in provided payload"
// @Failure 401 {string} json "Unauthorized"
// @Failure 500 {string} json "Internal Server Error"
func (h *Handler) Handler(c echo.Context) error {

	fmt.Println(c.Request())
	request := new(request.Product)

	_ = (&echo.DefaultBinder{}).Bind(request, c)

	fmt.Println("––--start request----")
	fmt.Println(request.Name)
	fmt.Println(request.Status)
	fmt.Println(request.Price)
	if errParam := c.Validate(request); errParam != nil {

		fmt.Println(&api.ErrorAPI{Message: fmt.Sprintf("Request Path %s", errParam.Error())})

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
