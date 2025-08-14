package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGoogleLogin(c *gin.Context) {
	url := AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	c.Status(http.StatusSeeOther)
	c.Redirect(http.StatusSeeOther, url)
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func HandleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != "randomstate" {
		c.String(http.StatusBadRequest, "states dont match")
		return
	}
	code := c.Query("code")
	googleconfig := GoogleConfig()

	token, err := googleconfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "code-token exchange failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "User data fetch failed")
		return
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON parsing of user data failed in server")
		return
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expires_in":    token.Expiry,
		"user_data":     json.RawMessage(userData),
	})
	// c.Data(http.StatusOK, "Application/json", userData)
}

func main() {
	// load the google oauth2 configs
	router := gin.Default()
	GoogleConfig()

	router.GET("/google_login", HandleGoogleLogin)
	router.GET("/google_callback", HandleGoogleCallback)

	log.Fatal(router.Run(":8080"))
}
