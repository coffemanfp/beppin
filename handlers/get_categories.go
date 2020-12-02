package handlers

import (
	"net/http"

	"github.com/coffemanfp/beppin/models"
	"github.com/labstack/echo"
)

// GetCategories - Get categories.
func GetCategories(c echo.Context) (err error) {
	var m models.ResponseMessage

	dbCategories, err := Storage.GetCategories()
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	if dbCategories == nil {
		dbCategories = make(models.Categories, 0)
	}

	m.Content = dbCategories
	m.ContentType = models.TypeCategories

	if m.Message == "" {
		m.Message = "Ok."
	}

	return c.JSON(http.StatusOK, m)
}
