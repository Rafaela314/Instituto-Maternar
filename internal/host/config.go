package host

import (
	"github.com/Rafaela314/instituto-maternar/models"
	"github.com/spf13/viper"
)

// NewAppConfig Retrieves the configuration
func NewAppConfig() (*models.AppConfig, error) {

	appConfig := &models.AppConfig{}

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		return nil, err
	}

	return appConfig, nil

}
