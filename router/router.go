package router

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/controllers"
	"github.com/labstack/echo"
)

// NewRouter - Creates the app router.
func NewRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/products", controllers.GetProducts)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
}
