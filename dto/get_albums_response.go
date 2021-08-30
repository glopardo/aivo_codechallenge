package dto

type (
	GetAlbumsResponse struct {
		AlbumItems []AlbumItem `json:"items"`
	}
	AlbumItem struct {
		Images      []Image `json:"images"`
		Name        string  `json:"name"`
		ReleaseDate string  `json:"release_date"`
		TotalTracks int     `json:"total_tracks"`
	}
	Image struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Url    string `json:"url"`
	}
)
