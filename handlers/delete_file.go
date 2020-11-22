package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
	"github.com/coffemanfp/beppin/utils"
	"github.com/labstack/echo"
)

// DeleteFile - Delete a file.
func DeleteFile(c echo.Context) (err error) {
	var m models.ResponseMessage
	var fileID int

	fileIDParam := c.Param("id")

	if fileID, err = utils.Atoi(fileIDParam); err != nil {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Get the old file info
	oldFile, err := Storage.GetFile(
		models.File{
			ID: int64(fileID),
		},
	)
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

	// Delete the system file
	err = os.Remove(oldFile.Path)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	id, err := Storage.DeleteFile(
		models.File{
			ID: int64(fileID),
		},
	)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Deleted."
	m.Content = models.File{
		ID: int64(id),
	}
	m.ContentType = models.TypeFile

	return c.JSON(http.StatusOK, m)
}
