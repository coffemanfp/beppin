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

// UpdateFile - Updates a file.
func UpdateFile(c echo.Context) (err error) {
	var m models.ResponseMessage

	fileIDParam := c.Param("id")
	fileID, err := utils.Atoi(fileIDParam)
	if err != nil || fileID == 0 {
		m.Error = fmt.Sprintf("%v: id", errs.ErrInvalidParam)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Get the new file info
	fileData, err := c.FormFile("file")
	if err != nil {
		m.Error = fmt.Sprintf("%v: file", errs.ErrInvalidData)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	var file models.File

	if err = c.Bind(&file); err != nil {
		m.Error = errs.ErrInvalidBody

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Set the new file path
	file.Path = models.NewFilePath(fileData.Filename)

	// Get old file info
	oldFile, err := Storage.GetFile(models.File{ID: int64(fileID)})
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

	// Save the new file
	err = utils.SaveMultipartFile(fileData, file.Path, 1000000)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidData) {
			m.Error = err.Error()

			return echo.NewHTTPError(http.StatusBadRequest, m)
		}
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusBadRequest, m)
	}

	// Delete the old file
	err = os.Remove(oldFile.Path)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	dbFile, err := Storage.UpdateFile(
		models.File{
			ID: int64(fileID),
		},
		file,
	)
	if err != nil {
		c.Logger().Error(err)
		m.Error = http.StatusText(http.StatusInternalServerError)

		return echo.NewHTTPError(http.StatusInternalServerError, m)
	}

	m.Message = "Updated."
	m.Content = dbFile
	m.ContentType = models.TypeFile
	return c.JSON(http.StatusOK, m)
}
