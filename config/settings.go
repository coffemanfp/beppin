package config

// Settings - Settings app.
type Settings struct {
	Port                     int    `json:"port" yaml:"port" mapstructure:"port"`
	LogsFile                 string `json:"logsFile" yaml:"logsFile" mapstructure:"logs_file"`
	MaxElementsPerPagination int    `json:"maxElementsPerPagination" yaml:"maxElementsPerPagination" mapstructure:"max_elements_per_pagination"`
	SecretKey                string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"`

	Database *Database `json:"database" yaml:"database"`
}

// Validate - Validates all settings.
func (s Settings) Validate() (valid bool) {
	valid = true

	if s.Database != nil {
		valid = s.Database.ValidateDatabase()
	} else {
		valid = false
	}

	switch "" {
	case s.LogsFile:
	case s.SecretKey:
		valid = false
	}

	switch 0 {
	case s.Port:
	case s.MaxElementsPerPagination:
		valid = false
	}

	return
}
