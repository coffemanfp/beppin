package config

import (
	"os"

	"github.com/spf13/viper"
)

// SetSettingsByEnv sets the env settings by environment vars.
func SetSettingsByEnv() (err error) {
	viper.SetEnvPrefix("beppin")
	viper.AutomaticEnv()

	viper.BindEnv(
		"port",
		"host",
		"assets",
		"logsFile",
		"secretKey",
		"maxElementsPerPagination",
		"maxImageSize",

		// Database
		"db_port",
		"db_name",
		"db_user",
		"db_password",
		"db_host",
		"db_sslMode",
		"db_url",
	)

	// Compatibility with others PaaS
	if viper.Get("port") == nil || viper.Get("port") == 8080 {
		viper.Set("port", os.Getenv("PORT"))
	}

	err = viper.Unmarshal(&GlobalSettings)
	if err != nil {
		return
	}

	err = viper.Unmarshal(&GlobalSettings.Database)
	return
}
