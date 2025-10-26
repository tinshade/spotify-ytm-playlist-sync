package main

import (
	"github.com/gin-gonic/gin"

	"tinshade.com/syps/api"
)

func main() {

	r := gin.Default()

	r.GET("/health", api.HealthCheck)
	r.GET("/token", api.GetSpotifyToken)
	r.GET("/get-playlist-songs", api.GetPlaylistSongs)
	r.Run(":8000")
}
