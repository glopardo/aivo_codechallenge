package service

import (
	"aivo-code-challenge/dto"
	"errors"
)

type (
	IntegrationRepository interface {
		GetAlbumsByArtistId(artistId, token string) ([]dto.Album, error)
		GetArtistID(artistName, accessToken string) (string, error)
		GetAccessToken() (string, error)
	}
	IntegrationService struct {
		repository IntegrationRepository
	}
)

func NewIntegrationService(repository IntegrationRepository) (IntegrationService, error) {
	if repository == nil {
		return IntegrationService{}, errors.New("repository is required but was empty")
	}
	return IntegrationService{repository: repository}, nil
}

func (s IntegrationService) GetAlbums(artistId, token string) ([]dto.Album, error) {
	return s.repository.GetAlbumsByArtistId(artistId, token)
}

func (s IntegrationService) GetArtist(artistName, token string) (string, error) {
	return s.repository.GetArtistID(artistName, token)
}

func (s IntegrationService) GetToken() (string, error) {
	return s.repository.GetAccessToken()
}
