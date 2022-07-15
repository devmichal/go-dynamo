package gethealthy

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Healthy is the handler for Healthy endpoint
// @Summary Check the health of the application
// @Router / [get]
// @Produce  json
// @Success 200 {string} json "Health check passed"
func Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK!")
}
