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

	// Users
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

}
