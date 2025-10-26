package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
)

func HealthCheck(c *gin.Context) {
	ResponseJSON(c, http.StatusOK, "Healthy and up!", nil)
}

func GetLoginToken() (*TokenResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     "your-client-id",
			"client_secret": "your-client-secret",
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
