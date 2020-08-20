package config

import "net/url"

// Settings - Settings app.
type Settings struct {
	Port                     int    `json:"port" yaml:"port" mapstructure:"port"`
	Host                     string `json:"host" yaml:"host" mapstructure:"host"`
	Assets                   string `json:"assets" yaml:"assets" mapstructure:"assets"`
	LogsFile                 string `json:"logsFile" yaml:"logsFile" mapstructure:"logs_file"`
	SecretKey                string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`
	Temps                    string `json:"temps" yaml:"temps" mapstructure:"temps"`
	MaxElementsPerPagination int    `json:"maxElementsPerPagination" yaml:"maxElementsPerPagination" mapstructure:"max_elements_per_pagination"`
	MaxImageSize             int64  `json:"maxImageSize" yaml:"maxImageSize" mapstructure:"max_image_size"`

	Database *Database `json:"database" yaml:"database"`
}

// Validate - Validates all settings.
func (s Settings) Validate() (valid bool) {
	valid = true

	if s.Database != nil {
		if !s.Database.ValidateDatabase() {
			valid = false
			return
		}
	} else {
		valid = false
		return
	}

	switch "" {
	case s.LogsFile:
	case s.SecretKey:
		valid = false
		return
	}

	switch 0 {
	case s.Port:
	case s.MaxElementsPerPagination:
		valid = false
	}

	if _, err := url.ParseRequestURI(s.Host); err != nil {
		valid = false
		return
	}

	if _, err := url.ParseRequestURI(s.Assets); err != nil {
		valid = false
		return
	}

	return
}

// ValidateMigrations - Validate only settings for migrations.
func (s Settings) ValidateMigrations() (valid bool) {
	valid = true

	if s.Database != nil {
		if !s.Database.ValidateDatabase() {
			valid = false
			return
		}
	} else {
		valid = false
		return
	}

	if s.LogsFile == "" {
		valid = false
	}
	return
}
