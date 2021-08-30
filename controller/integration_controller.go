package controller

import (
	"aivo-code-challenge/dto"
	"errors"

	"github.com/gin-gonic/gin"
)

type (
	IntegrationService interface {
		GetAlbums(artistId, token string) ([]dto.Album, error)
		GetArtist(artistName, token string) (string, error)
		GetToken() (string, error)
	}

	IntegrationController struct {
		service     IntegrationService
		accessToken string
	}
)

func NewIntegrationController(integrationService IntegrationService) (IntegrationController, error) {
	if integrationService == nil {
		return IntegrationController{}, errors.New("service is required but was empty")
	}
	accessToken, err := integrationService.GetToken()
	if err != nil {
		return IntegrationController{}, errors.New("error at getting access token")
	}
	return IntegrationController{
			service:     integrationService,
			accessToken: accessToken},
		nil
}

func (ctrl IntegrationController) GetBandDiscography(ctx *gin.Context) {
	var bandName = ctx.Param("q")
	artistId, err := ctrl.service.GetArtist(bandName, ctrl.accessToken)
	if err != nil {
		ctx.IndentedJSON(400, err.Error())
		return
	}

	albums, err := ctrl.service.GetAlbums(artistId, ctrl.accessToken)
	if err != nil {
		ctx.IndentedJSON(400, err.Error())
		return
	}
	ctx.IndentedJSON(200, albums)
}
