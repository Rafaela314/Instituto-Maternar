package host

import "github.com/Rafaela314/instituto-maternar/models"

// Host holds values and dependencies for handlers.
type Host struct {
	AppConfig *models.AppConfig
}

// NewHost creates a Host that shares variables and services for handlers.
func NewHost() *Host {

	host := &Host{}

	appConfig, err := NewAppConfig()
	if err != nil {
		panic(err)
	}
	host.AppConfig = appConfig

	return host

}
