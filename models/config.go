package models

type AppConfig struct {
	AppName         string `mapstructure:"APP_NAME"`
	AppCmdName      string // Application Cmd Name - Consul Service Name /AWS Url name.
	AppVersion      string `mapstructure:"APP_VERSION"` // Application Version (eg. "1.0.0")
	AppDescription  string
	CopyrightYears  string
	CopyrightHolder string
	DBName          string
	DBUser          string `mapstructure:"DB_USER"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          int    `mapstructure:"DB_PORT"`
	Host            string
	DNS             string
	Port            int
}
