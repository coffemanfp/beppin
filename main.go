package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})

	r.GET("/", controllers.GetProducts)

	r.Start(":8080")
}
