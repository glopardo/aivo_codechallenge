package server

import (
	"aivo-code-challenge/controller"
	"github.com/gin-gonic/gin"
)

type (
	Mapping struct {
		integrationController controller.IntegrationController
	}
)

func CreateMapping() Mapping {
	return Mapping{
		integrationController: resolveIntegrationController(),
	}
}

func (m Mapping) mapUrlsToControllers(router *gin.Engine) {
	m.controllerRoutes(router)
}

func (m Mapping) controllerRoutes(router *gin.Engine) {
	router.GET("/api/v1/albums/:q", m.integrationController.GetBandDiscography)
}
