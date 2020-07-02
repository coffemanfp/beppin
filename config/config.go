package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/coffemanfp/beppin-server/utils"
	yaml "gopkg.in/yaml.v3"
)

var settings *Settings

// Settings - Settings app.
type Settings struct {
	Port                     int       `json:"port" yaml:"port"`
	Database                 *Database `json:"database" yaml:"database"`
	MaxElementsPerPagination int       `json:"maxElementsPerPagination" yaml:"maxElementsPerPagination"`
}

// Database - Database settings.
type Database struct {
	Name     string `json:"name" yaml:"name"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
}

// GetSettings - Get the server settings.
//	@return s *Settings:
//		Server settings.
func GetSettings() (s Settings, err error) {
	if settings == nil {
		err = fmt.Errorf("error settings not found")
		return
	}

	s = *settings

	return
}

// SetSettingsByFile - Sets the settings by a file.
//	@param filePath string:
//		Config filepath.
func SetSettingsByFile(filePath string) (err error) {
	fileBytes, err := utils.GetFilebytes(filePath)
	if err != nil {
		return
	}

	switch ext := filepath.Ext(filePath)[1:]; {
	case ext == "json":
		err = json.Unmarshal(fileBytes, &settings)
	case ext == "yaml":
		err = yaml.Unmarshal(fileBytes, &settings)
	default:
		err = fmt.Errorf("extension (%s) not supported:\n%s", ext, err)
		if err != nil {
			return
		}
	}
	if err != nil {
		err = fmt.Errorf("failed to unmarshalling the settings:\n%s", err)
		return
	}

	fmt.Println(os.Getenv("PORT"))

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		var port int
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			err = fmt.Errorf("failed to get the port environment variable:\n%s", err)
			return
		}

		settings.Port = port
	}

	if !settings.Validate() {
		err = errors.New("settings are not populated")
		return
	}

	return
}

// Validate - Validates all settings.
func (s Settings) Validate() (valid bool) {
	valid = true

	if s.Port == 0 {
		valid = false
	}

	valid = s.ValidateDatabase()
	return
}

// ValidateDatabase - Validates the database settings.
func (s Settings) ValidateDatabase() (valid bool) {
	valid = true

	switch "" {
	case s.Database.Host:
	case s.Database.Name:
	case s.Database.User:
	case s.Database.Password:
		valid = false
	}
	if s.Database.Port == 0 {
		valid = false
	}
	return
}
