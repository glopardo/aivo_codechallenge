package dto

type (
	Album struct {
		Name     string     `json:"name"`
		Released string     `json:"released"`
		Tracks   int        `json:"tracks"`
		Cover    AlbumCover `json:"cover"`
	}

	AlbumCover struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Url    string `json:"url"`
	}
)
