package config

import "github.com/spf13/viper"

// ValidateSettings validates settings for a specific action.
func ValidateSettings(action string) (valid bool) {
	switch action {
	case "database":
		valid = validateDatabase()
	case "migrations":
		valid = validateMigrations()
	}
	return
}

func validateDatabase() (valid bool) {
	valid = true

	if viper.GetString("db_name") == "" ||
		viper.GetInt("db_port") == 0 ||
		viper.GetString("db_user") == "" ||
		viper.GetString("db_password") == "" ||
		viper.GetString("db_host") == "" ||
		viper.GetString("db_sslMode") == "" ||
		viper.GetString("db_url") == "" {
		valid = false
	}
	return
}

func validateMigrations() (valid bool) {
	valid = true

	if !validateDatabase() {
		valid = false
		return
	}

	if viper.GetString("logsFile") == "" {
		valid = false
	}
	return
}
