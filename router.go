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
	r.POST("/login/:provider", controllers.LoginWithProvider)

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

	// Products
	r.POST("/products", controllers.CreateProduct, middleware.JWTWithConfig(jwtConfig))
	r.PUT("/products/:id", controllers.UpdateProduct, middleware.JWTWithConfig(jwtConfig))
	r.DELETE("/products/:id", controllers.DeleteProduct, middleware.JWTWithConfig(jwtConfig))

	// Users
	r.GET("/users", controllers.GetUsers, middleware.JWTWithConfig(jwtConfig))
	r.GET("/users/:id", controllers.GetUser, middleware.JWTWithConfig(jwtConfig))
	r.PUT("/users/:id", controllers.UpdateUser, middleware.JWTWithConfig(jwtConfig))
	r.DELETE("/users/:id", controllers.DeleteUser, middleware.JWTWithConfig(jwtConfig))

	return
}
