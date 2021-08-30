package server

import (
	"aivo-code-challenge/controller"
	"aivo-code-challenge/repository"
	"aivo-code-challenge/service"
)

func resolveIntegrationRepository() repository.IntegrationRepository {
	dao, _ := repository.NewIntegrationRepository()
	return dao
}

func resolveIntegrationService() service.IntegrationService {
	repository := resolveIntegrationRepository()
	service, err := service.NewIntegrationService(repository)
	if err != nil {
		panic("PANIC on resolveIntegrationService")
	}
	return service
}

func resolveIntegrationController() controller.IntegrationController {
	integrationService := resolveIntegrationService()
	ctrl, err := controller.NewIntegrationController(integrationService)
	if err != nil {
		panic("PANIC on resolveIntegrationController")
	}
	return ctrl
}
