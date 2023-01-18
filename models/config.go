package models

type AppConfig struct {
	AppName         string `mapstructure:"APP_NAME"`
	AppCmdName      string // Application Cmd Name - Consul Service Name /AWS Url name.
	AppVersion      string `mapstructure:"APP_VERSION"` // Application Version (eg. "1.0.0")
	AppDescription  string
	CopyrightYears  string
	CopyrightHolder string
	DBName          string `mapstructure:"POSTGRES_NAME"`
	DBUser          string `mapstructure:"POSTGRES_USER"`
	DBHost          string `mapstructure:"POSTGRES_HOST"`
	DBPort          int    `mapstructure:"POSTGRES_PORT"`
	DBPassword      string `mapstructure:"POSTGRES_PASSWORD"`
	Host            string
	DNS             string
	Port            int
}
