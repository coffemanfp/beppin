package config

// GlobalSettings are the global settings.
var GlobalSettings *Settings

// Settings represents the global settings.
type Settings struct {
	Port                     int    `json:"port" yaml:"port" mapstructure:"port"`
	Host                     string `json:"host" yaml:"host" mapstructure:"host"`
	Assets                   string `json:"assets" yaml:"assets" mapstructure:"assets"`
	LogsFile                 string `json:"logsFile" yaml:"logsFile" mapstructure:"logsFile"`
	SecretKey                string `json:"secret_key" yaml:"secret_key" mapstructure:"secretKey"`
	Temps                    string `json:"temps" yaml:"temps" mapstructure:"temps"`
	MaxElementsPerPagination int    `json:"maxElementsPerPagination" yaml:"maxElementsPerPagination" mapstructure:"maxElementsPerPagination"`
	MaxImageSize             int    `json:"maxImageSize" yaml:"maxImageSize" mapstructure:"maxImageSize"`

	Database *Database `json:"database" yaml:"database"`
}

// Database represents the database settings.
type Database struct {
	Name     string `json:"name" yaml:"name" mapstructure:"db_name"`
	Port     int    `json:"port" yaml:"port" mapstructure:"db_port"`
	User     string `json:"user" yaml:"user" mapstructure:"db_user"`
	Password string `json:"password" yaml:"password" mapstructure:"db_password"`
	Host     string `json:"host" yaml:"host" mapstructure:"db_host"`
	SslMode  string `json:"sslMode" yaml:"db_ssl_mode" mapstructure:"db_sslmode"`
	URL      string `json:"url" yaml:"url" mapstructure:"db_url"`
}
