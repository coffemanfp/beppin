package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/coffemanfp/beppin-server/utils"
	"github.com/lib/pq"
	"github.com/spf13/viper"
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

// GetSettings - Get the server settings.
//	@return s *Settings:
//		Server settings.
func GetSettings() (s Settings) {
	if settings == nil {
		setDefaultSettings()
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
	viper.SetEnvPrefix("beppin")

	err = bindEnvVars()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(settings)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal settings:\n%s", err)
		return
	}

	var databaseURL string
	if settings.Database.URL == "" {
		databaseURL, err = settings.Database.GetURL()
		if err != nil {
			return
		}
	} else {
		var databaseURL string
		databaseURL, err = pq.ParseURL(settings.Database.URL)
		if err != nil {
			err = fmt.Errorf("failed to parse the database url connection:\n%s", err)
			return
		}

		databaseURL += " sslmode=" + settings.Database.SslMode
	}

	settings.Database.URL = databaseURL
	return
}

// bindEnvVars binds the environment variables.
// returns an error with a list of the missing variables.
func bindEnvVars() (err error) {
	envVarNames := []string{
		"port", "logs_file",
		"max_elements_per_pagination", "secret_key",
		"db_name", "db_user",
		"db_port", "db_password",
		"db_host", "db_sslmode", "db_url",
	}

	var missingVariables []error
	for _, envVarName := range envVarNames {
		err = viper.BindEnv(envVarName)
		if err != nil {
			missingVariables = append(missingVariables, err)
		}
	}

	for _, missingVariable := range missingVariables {
		err = errors.New(missingVariable.Error() + "\n")
	}

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

	valid = s.Database.ValidateDatabase()
	return
}
