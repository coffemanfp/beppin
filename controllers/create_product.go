package controllers

import (
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
)

// CreateProduct - Creates a product.
func CreateProduct(c echo.Context) (err error) {
	var m models.ResponseMessage
	var product models.Product

	if err = c.Bind(&product); err != nil {
		m.Error = "invalid body"

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	db, err := database.Get()
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	err = dbu.InsertProduct(db, product)
	if err != nil {
		c.Logger().Error(err)

		return echo.ErrInternalServerError
	}

	m.Message = "Created."

	return c.JSON(http.StatusCreated, m)
}
