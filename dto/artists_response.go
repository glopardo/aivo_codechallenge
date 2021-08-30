package dto

type (
	ArtistResponse struct {
		Body ArtistsBody `json:"artists"`
	}
	ArtistsBody struct {
		Items []ArtistItem `json:"items"`
	}
	ArtistItem struct {
		Id string `json:"id"`
	}
)
