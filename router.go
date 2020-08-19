package main

import (
	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/controllers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter(e *echo.Echo) {
	settings := config.GetSettings()

	// API group
	r := e.Group("/v1")

	// Sign In
	r.POST("/login", controllers.Login)

	// Sign Up
	r.POST("/signup", controllers.SignUp)

	// Products
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)

	// JWT Middleware
	jwtConfig := middleware.JWTConfig{
		Claims:      &models.Claim{},
		SigningKey:  []byte(settings.SecretKey),
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}

	jwtMiddleware := middleware.JWTWithConfig(jwtConfig)

	// Products
	r.POST("/products", controllers.CreateProduct, jwtMiddleware)
	r.PUT("/products/:id", controllers.UpdateProduct, jwtMiddleware)
	r.DELETE("/products/:id", controllers.DeleteProduct, jwtMiddleware)

	// Users
	r.GET("/users", controllers.GetUsers, jwtMiddleware)
	r.GET("/users/:id", controllers.GetUser, jwtMiddleware)
	r.PUT("/users/:id", controllers.UpdateUser, jwtMiddleware)
	r.PUT("/users/:id/avatar", controllers.UpdateAvatar, jwtMiddleware)
	r.DELETE("/users/:id", controllers.DeleteUser, jwtMiddleware)

	return
}
