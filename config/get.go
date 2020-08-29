package config

import "github.com/spf13/viper"

// GetVar gets the specified env config var.
func GetVar(name string) (value interface{}) {
	switch name {
	case "port":
		value = viper.GetInt("port")
	case "host":
		value = viper.GetString("host")
	case "assets":
		value = viper.GetString("assets")
	case "logsFile":
		value = viper.GetString("logsFile")
	case "secretKey":
		value = viper.GetString("secretKey")
	case "maxElementsPerPagination":
		value = viper.GetInt("maxElementsPerPagination")
	case "maxImageSize":
		value = viper.GetInt("maxImageSize")

	// Database
	case "db_port":
		value = viper.GetInt("db_port")
	case "db_name":
		value = viper.GetString("db_name")
	case "db_user":
		value = viper.GetString("db_user")
	case "db_password":
		value = viper.GetString("db_password")
	case "db_host":
		value = viper.GetString("db_host")
	case "db_sslMode":
		value = viper.GetString("db_sslMode")
	}
	return
}
