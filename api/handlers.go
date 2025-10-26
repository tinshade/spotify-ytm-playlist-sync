package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	ResponseJSON(c, http.StatusOK, "Healthy and up!", nil)
}

func GetPlaylistSongs(c *gin.Context) {
	data, err := FetchPlaylistTracks("6gUTRutxUlCS2RGV0otHeo", 10, 0) //TODO: Get from user
	if err != nil {
		fmt.Print(err)
		ResponseJSON(c, 500, "Something went wrong while getting tracks from the playlist", data)
		return
	}
	ResponseJSON(c, http.StatusOK, "Playlist tracks fetched successfully!", data)
}

func GetSpotifyToken(c *gin.Context) {
	token, err := FetchSpotifyToken()
	if err != nil {
		fmt.Print(err)
		ResponseJSON(c, 500, "Something went wrong while getting Spotify Login Token", token)
		return
	}
	ResponseJSON(c, http.StatusOK, "Token fetched successfully!", token)
}
