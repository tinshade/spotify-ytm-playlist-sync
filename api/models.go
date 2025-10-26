package api

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SimplifiedTrack struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Href   string `json:"href"`
	SongID string `json:"id"`
}

type PlaylistItem struct {
	Track SimplifiedTrack `json:"track"`
}

type PlaylistTracks struct {
	Href     string         `json:"href"`
	Limit    int            `json:"limit"`
	Offset   int            `json:"offset"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Total    int            `json:"total"`
	Items    []PlaylistItem `json:"items"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}
