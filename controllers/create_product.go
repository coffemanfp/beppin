package controllers

import (
	"errors"
	"net/http"

	"github.com/coffemanfp/beppin-server/database"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	"github.com/coffemanfp/beppin-server/models"
	"github.com/labstack/echo"
)

// CreateProduct - Create a product.
func CreateProduct(c echo.Context) (err error) {
	m := models.ResponseMessage{}
	product := models.Product{}

	if err = c.Bind(&product); err != nil {
		m.Error = errors.New("invalid body")

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

	return c.String(http.StatusCreated, "")
}
