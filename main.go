package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/coffemanfp/beppin-server/controllers"
	"github.com/coffemanfp/beppin-server/database"
	"github.com/labstack/echo"
)

func main() {
	err := config.SetSettingsByFile("config.yaml")
	if err != nil {
		log.Fatal("failed to load settings:\n%s", err)
	}

	settings, err := config.GetSettings()
	if err != nil {
		log.Fatal("failed to get settings:\n%s", err)
	}

	_, err = database.OpenConn()
	if err != nil {
		log.Fatal("failed to init database:\n%s", err)
	}

	r := echo.New()

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})

	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	log.Println(r.Start(fmt.Sprintf(":%d", settings.Port)))
}
