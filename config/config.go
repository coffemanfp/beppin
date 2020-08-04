package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/coffemanfp/beppin-server/utils"
	"github.com/lib/pq"
	yaml "gopkg.in/yaml.v3"
)

var settings *Settings

// Settings - Settings app.
type Settings struct {
	Port                     int    `json:"port" yaml:"port"`
	LogsFile                 string `json:"logsFile" yaml:"logsFile"`
	MaxElementsPerPagination int    `json:"maxElementsPerPagination" yaml:"maxElementsPerPagination"`

	SecretKey string    `json:"secretKey" yaml:"secretKey"`
	Database  *Database `json:"database" yaml:"database"`
}

// Database - Database settings.
type Database struct {
	Name     string `json:"name" yaml:"name"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	SslMode  string `json:"sslMode" yaml:"sslMode"`
	URL      string `json:"url" yaml:"url"`
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

	if !settings.Validate() {
		err = errors.New("settings are not populated")
		return
	}

	return
}

// SetSettingsByEnv - Fills the settings by the environment variables
func SetSettingsByEnv() (err error) {
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		var port int
		port, err = strconv.Atoi(portEnv)
		if err != nil {
			err = fmt.Errorf("failed to get the port environment variable:\n%s", err)
			return
		}

		settings.Port = port
	}

	var databaseURL string
	if databaseURLEnv := os.Getenv("DATABASE_URL"); databaseURLEnv != "" {
		databaseURL, err = pq.ParseURL(databaseURLEnv)
		if err != nil {
			err = fmt.Errorf("failed to parse the database url connection:\n%s", err)
			return
		}
		databaseURL += " sslmode=" + settings.Database.SslMode

	} else {
		databaseURL = fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
			settings.Database.User,
			settings.Database.Password,
			settings.Database.Name,
			settings.Database.Host,
			settings.Database.Port,
			settings.Database.SslMode,
		)
	}
	settings.Database.URL = databaseURL
	return
}

func setDefaultSettings() {
	settings = &Settings{
		Port:                     8080,
		LogsFile:                 "logs/server.log",
		MaxElementsPerPagination: 20,
		SecretKey:                "Security",

		Database: &Database{
			Port:     5432,
			Name:     "database_name",
			User:     "database_user",
			Password: "database_password",
			Host:     "localhost",
			SslMode:  "disabled",
		},
	}
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
