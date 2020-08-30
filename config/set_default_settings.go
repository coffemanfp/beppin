package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SetDefaultSettings sets the default settings.
func SetDefaultSettings() (err error) {
	viper.SetDefault("port", 8080)
	viper.SetDefault(
		"host",
		fmt.Sprintf(
			"http://%s:%d",
			"localhost",
			viper.GetInt("port"),
		),
	)
	viper.SetDefault("assets", "assets/")
	viper.SetDefault("logsFile", "logs/server.log")
	viper.SetDefault("secretKey", "Security")
	viper.SetDefault("maxElementsPerPagination", 20)
	viper.SetDefault("maxImageSize", 3e+6)

	// Database
	viper.SetDefault("db_port", 5432)
	viper.SetDefault("db_name", "beppin_tests")
	viper.SetDefault("db_user", "beppin_tests")
	viper.SetDefault("db_password", "beppin_tests")
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_sslMode", "disable")
	viper.SetDefault(
		"db_url",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
			viper.GetString("db_user"),
			viper.GetString("db_password"),
			viper.GetString("db_name"),
			viper.GetString("db_host"),
			viper.GetInt("db_port"),
			viper.GetString("db_sslMode"),
		),
	)

	err = viper.Unmarshal(&GlobalSettings)
	if err != nil {
		return
	}

	err = viper.Unmarshal(&GlobalSettings.Database)
	return
}
