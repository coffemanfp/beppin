package config

import (
	"fmt"
	"path/filepath"
	"strings"

	errs "github.com/coffemanfp/beppin-server/errors"

	"github.com/lib/pq"
	"github.com/spf13/viper"
)

var settings *Settings

// GetSettings - Get the server settings.
//	@return s Settings:
//		Server settings.
func GetSettings() (s Settings) {
	if settings == nil {
		SetDefaultSettings()
	}

	s = *settings
	return
}

// SetDefaultSettings configure the default settings values.
func SetDefaultSettings() {
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
			SslMode:  "disable",
		},
	}
}

// SetSettingsByFile - Sets the settings by a file.
//	@param path string:
//		Config filepath.
func SetSettingsByFile(path string) (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType(filepath.Ext(path)[1:])
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	err = viper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("failed to read in config: %w", err)
		return
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal settings: %v", err)
		return
	}

	if !settings.Validate() {
		err = fmt.Errorf("failed to validate settings: %w", errs.ErrInvalidSettings)
	}
	return
}

// SetSettingsByEnv - Fills the settings by the environment variables
func SetSettingsByEnv() (err error) {
	viper.SetEnvPrefix("beppin")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = bindEnvVars()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(&settings)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal settings: %v", err)
		return
	}

	err = viper.Unmarshal(&settings.Database)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal database settings: %v", err)
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
			err = fmt.Errorf("failed to parse database url: %v", err)
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
		"db_host", "db_ssl_mode", "db_url",
	}

	var missingVariables []error
	for _, envVarName := range envVarNames {
		err = viper.BindEnv(envVarName)
		if err != nil {
			missingVariables = append(missingVariables, err)
		}
	}

	for _, missingVariable := range missingVariables {
		err = fmt.Errorf("failed to get env var: %v", missingVariable)
	}

	return
}
