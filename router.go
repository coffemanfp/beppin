package main

import (
	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/handlers"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter(e *echo.Echo) {

	// API group
	r := e.Group("/v1")

	// Sign In
	r.POST("/login", handlers.Login)
	r.POST("/login/:provider", handlers.LoginWithProvider)

	// Sign Up
	r.POST("/signup", handlers.SignUp)

	// Products
	r.GET("/products", handlers.GetProducts)
	r.GET("/products/:id", handlers.GetProduct)

	// JWT Middleware
	jwtConfig := middleware.JWTConfig{
		Claims:      &models.Claim{},
		SigningKey:  []byte(config.GetVar("secretKey").(string)),
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}

	jwtMiddleware := middleware.JWTWithConfig(jwtConfig)

	// Products
	r.POST("/products", handlers.CreateProduct, jwtMiddleware)
	r.PUT("/products/:id", handlers.UpdateProduct, jwtMiddleware)
	r.DELETE("/products/:id", handlers.DeleteProduct, jwtMiddleware)

	// Users
	r.GET("/users", handlers.GetUsers, jwtMiddleware)
	r.GET("/users/:id", handlers.GetUser, jwtMiddleware)
	r.PUT("/users/:id", handlers.UpdateUser, jwtMiddleware)
	r.PUT("/users/:id/avatar", handlers.UpdateAvatar, jwtMiddleware)
	r.DELETE("/users/:id", handlers.DeleteUser, jwtMiddleware)

	return
}
