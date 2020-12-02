package handlers

import (
	"errors"
	"fmt"
	"net/http"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/labstack/echo"
)

// UploadFile - Handles the files upload.
func UploadFile(c echo.Context) (err error) {
	var m models.ResponseMessage

	fileData, err := c.FormFile("file")
	if err != nil {
		m.Error = fmt.Sprintf("%v: file", errs.ErrInvalidData)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	destination := models.NewFilePath(fileData.Filename)

	// Saving file on the filesystem
	err = utils.SaveMultipartFile(fileData, destination, 1000000)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidData) {
			m.Error = err.Error()

			return echo.NewHTTPError(http.StatusBadRequest, m)
		}
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	file := models.File{
		Path: destination,
	}

	dbFile, err := Storage.CreateFile(file)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Created."
	m.Content = dbFile
	m.ContentType = models.TypeFile

	return c.JSON(http.StatusCreated, m)
}
