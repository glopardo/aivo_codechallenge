package repository

import (
	"aivo-code-challenge/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type (
	IntegrationRepository struct {
	}
)

const (
	httpRespStatusCodeSuccess = "200"
	APIAuthUrl                = "https://accounts.spotify.com"
	APIURL                    = "https://api.spotify.com"

	// TODO secrets
	apiAuthorization = "Basic NWUwZDFmNDkxOTFkNDE5MDk4NWQ4MWVjMWFmNzk5YzI6MDc3MzY5MGMxYzhlNDk0ZmI3YmNhN2E0Mzc2NmIyYjU="
)

func NewIntegrationRepository() (IntegrationRepository, error) {
	return IntegrationRepository{}, nil
}

func (repo IntegrationRepository) GetAlbumsByArtistId(artistId, token string) ([]dto.Album, error) {
	resource := fmt.Sprintf("/v1/artists/%s/albums", artistId)
	data := url.Values{}

	u, _ := url.ParseRequestURI(APIURL)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(data.Encode()))
	fmt.Println("token: ", token)
	q := r.URL.Query()
	r.URL.RawQuery = q.Encode()
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", token))

	resp, _ := client.Do(r)

	if resp.Status != httpRespStatusCodeSuccess {
		return nil, errors.New("couldn't find any album for this artist")
	}
	body := &dto.GetAlbumsResponse{}
	_ = json.NewDecoder(resp.Body).Decode(body)
	return mapAlbumItemsToDTO(body.AlbumItems), nil
}

func mapAlbumItemsToDTO(albumItems []dto.AlbumItem) []dto.Album {
	retAlbums := []dto.Album{}
	for _, albumItem := range albumItems {
		retAlbums = append(retAlbums, dto.Album{Name: albumItem.Name, Released: albumItem.ReleaseDate, Tracks: albumItem.TotalTracks, Cover: dto.AlbumCover{Url: albumItem.Images[0].Url, Height: albumItem.Images[0].Height, Width: albumItem.Images[0].Width}})
	}
	return retAlbums
}

func (repo IntegrationRepository) GetArtistID(artistName, accessToken string) (string, error) {
	resource := "/v1/search/"
	data := url.Values{}

	u, _ := url.ParseRequestURI(APIURL)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, strings.NewReader(data.Encode()))
	fmt.Println("token: ", accessToken)
	q := r.URL.Query()
	q.Add("q", artistName)
	q.Add("type", "artist")
	r.URL.RawQuery = q.Encode()
	r.Header.Add("Authorization", fmt.Sprintf("%s %s", "Bearer", accessToken))

	resp, _ := client.Do(r)
	body := &dto.ArtistResponse{}
	_ = json.NewDecoder(resp.Body).Decode(body)
	if resp.Status != httpRespStatusCodeSuccess {
		return "", errors.New(fmt.Sprintf("couldn't find %s", artistName))
	}
	// Se asume que el primer resultado que viene es la mejor ocurrencia
	// por ende es el que se devuelve como resultado válido
	return body.Body.Items[0].Id, nil
}

func (repo IntegrationRepository) GetAccessToken() (string, error) {
	resource := "/api/token/"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	u, _ := url.ParseRequestURI(APIAuthUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Authorization", apiAuthorization)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	if resp.Status != httpRespStatusCodeSuccess {
		return "", errors.New("error getting access token")
	}
	body := &dto.AccessTokenResponse{}
	_ = json.NewDecoder(resp.Body).Decode(body)
	return body.AccessToken, nil
}
