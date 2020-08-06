package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/coffemanfp/beppin-server/utils"
	"github.com/labstack/echo"
)

// NewLogger - Sets a new logger.
func NewLogger(e *echo.Echo, path string) (err error) {
	var logFile *os.File

	err = os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		err = fmt.Errorf("failed to create %s folder: %v", path, err)
		return
	}

	exist, err := utils.ExistsFile(path)
	if err != nil {
		return
	}

	if !exist {
		logFile, err = os.Create(path)
		if err != nil {
			err = fmt.Errorf("failed to create %s file: %v", path, err)
			return
		}
	} else {
		logFile, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			err = fmt.Errorf("failed to open %s file: %v", path, err)
			return
		}
	}

	mw := io.MultiWriter(e.Logger.Output(), logFile)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e.Logger.SetOutput(mw)
	return
}
