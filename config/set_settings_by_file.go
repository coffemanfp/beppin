package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SetSettingsByFile sets the env settings by config files.
func SetSettingsByFile() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/beppin")
	viper.AddConfigPath("$HOME/.beppin")

	err = viper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("failed to read the config file: \n%v", err)
	}
	return
}
