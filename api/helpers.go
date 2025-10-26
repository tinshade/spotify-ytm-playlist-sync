package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const SpotifyBaseAPIURL string = "https://api.spotify.com"
const Version string = "1"

func FetchSpotifyToken() (*TokenResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     os.Getenv("SPOTIFY_CLIENT_ID"),
			"client_secret": os.Getenv("SPOTIFY_CLIENT_SECRET"),
		}).
		Post("https://accounts.spotify.com/api/token")

	if err != nil {
		return nil, err
	}

	var token TokenResponse

	// Convert response into JSON
	if err := json.Unmarshal(resp.Body(), &token); err != nil {
		return nil, err
	}

	return &token, nil

}

func generatePlaylistFetchURL(playlistID string, limit int, offset int) string {
	const base = "%s/v%s/playlists/%s/tracks?limit=%d&offset=%d"
	formattedURL := fmt.Sprintf(base, SpotifyBaseAPIURL, Version, playlistID, limit, offset)
	return formattedURL
}

func getSimlifiedTracksDetails(data PlaylistTracks) ([]SimplifiedTrack, error) {
	var simplifiedTracks []SimplifiedTrack
	for _, item := range data.Items {
		simplifiedTracks = append(simplifiedTracks, item.Track)
	}

	return simplifiedTracks, nil
}

func FetchPlaylistTracks(playlistID string, limit int, offset int) ([]SimplifiedTrack, error) {
	token, err := FetchSpotifyToken()
	if err != nil {
		return nil, err
	}
	formattedURL := generatePlaylistFetchURL(playlistID, limit, offset)
	client := resty.New()
	resp, err := client.R().SetHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken)).Get(formattedURL)

	if err != nil {
		return nil, err
	}

	var data PlaylistTracks
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	simplifiedTracks, err := getSimlifiedTracksDetails(data)

	return simplifiedTracks, err
}
