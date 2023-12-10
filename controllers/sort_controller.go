package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"backend/models"
	"backend/services"
)

func ProcessSortSequential(c echo.Context) error {
	var request models.SortRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	response := services.ProcessSortSequential(request)

	return c.JSON(http.StatusOK, response)
}

func ProcessSortConcurrent(c echo.Context) error {
	var request models.SortRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON payload"})
	}

	response := services.ProcessSortConcurrent(request)

	return c.JSON(http.StatusOK,response)
}