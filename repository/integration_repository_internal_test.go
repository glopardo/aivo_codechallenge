package repository

import (
	"aivo-code-challenge/dto"
	"testing"
)

func TestNewIntegrationRepository(t *testing.T) {
	_, err := NewIntegrationRepository()
	assertEqual(t, err, nil)
}

func TestMapAlbumItemsToDTO_success(t *testing.T) {
	want := dto.Album{
		Name:     "Nombre",
		Released: "Lanzamiento",
		Tracks:   11,
		Cover:    dto.AlbumCover{
			Height: 150,
			Width:  150,
			Url:    "URL",
		},
	}
	albumItem := []dto.AlbumItem{
		{
			Images:		 []dto.Image{{
				Height: 150,
				Width:  150,
				Url:    "URL",
			}},
			Name:        "Nombre",
			ReleaseDate: "Lanzamiento",
			TotalTracks: 11,
		},
	}
	albums := mapAlbumItemsToDTO(albumItem)
	assertEqual(t, albums[0], want)
}

func assertEqual(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Fatalf("%s != %s", got, want)
	}
}