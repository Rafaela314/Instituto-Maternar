package host

import (
	"github.com/Rafaela314/instituto-maternar/models"
	"github.com/gin-gonic/gin"
)

const (
	// APIV1 allows us to specify a versioning.
	APIV1 = "/v1/"
)

// Host holds values and dependencies for handlers.
type Host struct {
	AppConfig *models.AppConfig
	Engine    *gin.Engine
}

// NewHost creates a Host that shares variables and services for handlers.
func NewHost() *Host {

	host := &Host{}

	appConfig, err := NewAppConfig()
	if err != nil {
		panic(err)
	}
	host.AppConfig = appConfig

	host.Engine = host.CreateWebRouter()

	return host

}
