// @title commission-service
// @version 1.0
// @description microservice of commission-service
// @BasePath /
package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"musement.com/commission-service/internal/api/calculate"
	"musement.com/commission-service/internal/api/getcalculate"
	"musement.com/commission-service/internal/api/gethealthy"
	"musement.com/commission-service/internal/database"
	"musement.com/commission-service/internal/repository"
	"musement.com/commission-service/internal/validation"
)

const (
	serviceName    = "commission-service"
	serviceVersion = "v1"
)

func Command(c *cli.Context) (err error) {
	// Echo instance
	e := echo.New()
	e.Validator = validation.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"*"},
	}))

	// Dependencies
	orm := database.ConnectDynamo()

	commissionRepository := repository.NewRuleRepository(orm)

	postCalculator := calculate.New(*commissionRepository)
	getCalculate := getcalculate.New(*commissionRepository)

	// Route => handler
	e.GET("/", gethealthy.Handle)
	e.POST("/v1/calculate", postCalculator.Handler)
	e.GET("/v1/calculate/:id", getCalculate.Handle)

	// Start server
	fmt.Println(e.Start(":8010"))

	return nil
}
