package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// CreateProduct - Creates a product.
func CreateProduct(c echo.Context) (err error) {
	var m models.ResponseMessage
	var product models.Product

	userIDToken := c.Get("user").(*jwt.Token).Claims.(*models.Claim).User.ID

	if err = c.Bind(&product); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	product.UserID = userIDToken

	if !product.Validate() {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Check if exists the files
	if len(product.Images) != 0 {
		for _, image := range product.Images {
			// Get old file info
			oldFile, err := Storage.GetFile(models.File{ID: int64(image.ID)})
			if err != nil {
				if errors.Is(err, errs.ErrNotExistentObject) {
					m.Error = fmt.Sprintf("%v: file", errs.ErrNotExistentObject)

					return echo.NewHTTPError(http.StatusNotFound, m)
				}
				c.Logger().Error(err)
				m.Error = http.StatusText(http.StatusInternalServerError)

				return echo.NewHTTPError(http.StatusInternalServerError, m)
			}

			// Check if the file exists in the filesystem
			exists, err := utils.ExistsFile(oldFile.Path)
			if err != nil {
				c.Logger().Error(err)
				m.Error = http.StatusText(http.StatusInternalServerError)

				return echo.NewHTTPError(http.StatusInternalServerError, m)
			}

			if !exists {
				m.Error = fmt.Sprintf("%v: file", errs.ErrNotExistentObject)

				return echo.NewHTTPError(http.StatusNotFound, m)
			}

		}
	}

	createdProduct, err := Storage.CreateProduct(product)
	if err != nil {
		if errors.Is(err, errs.ErrNotExistentObject) {
			m.Error = fmt.Sprintf("%v: user", errs.ErrNotExistentObject)

			return echo.NewHTTPError(http.StatusNotFound, m)
		}
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Created."
	m.Content = createdProduct
	m.ContentType = models.TypeProduct
	return c.JSON(http.StatusCreated, m)
}
