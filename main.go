package main

import (
	"github.com/labstack/echo/v4"
	"backend/controllers"
	
)

func main() {
	e := echo.New()

	e.POST("/process-single", controllers.ProcessSortSequential)
	e.POST("/process-concurrent", controllers.ProcessSortConcurrent)

	e.Logger.Fatal(e.Start(":8000"))
}