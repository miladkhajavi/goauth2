package controller

import (
	"context"
	"fmt"
	"goauth/config"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"response": "welcome",
	// })
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	// log.Fatal("-------------", url)
	c.Redirect(http.StatusFound, url)
}

func GoogleCallback(c *gin.Context) {

	state := c.Request.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Printf("doesnt match state")
		return
	}
	//code
	code := c.Request.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Print(err)
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Print(err)
	}
	//parse response
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": userData,
		"test":     resp.Body,
	})
}
