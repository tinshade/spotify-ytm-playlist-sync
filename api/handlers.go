package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	ResponseJSON(c, http.StatusOK, "Healthy and up!", nil)
}

func GetPlaylistSongs(c *gin.Context) {
	var playlistID string = c.Param("playlistID") //6gUTRutxUlCS2RGV0otHeo
	fmt.Println("Playlist ID:", playlistID)
	if playlistID == "" {
		ResponseJSON(c, 400, "Playlist ID is required", nil)
		return
	}

	var limitString string = c.DefaultQuery("limit", "10")
	var offsetString string = c.DefaultQuery("offset", "0")

	// Convert to int or use default values
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		offset = 0
	}

	data, err := FetchPlaylistTracks(playlistID, limit, offset)
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
