package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

var todos []string
var toValidate = map[string]string{
	"aud": "api://default",
	"cid": os.Getenv("OKTA_CLIENT_ID"),
}

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func ListItem(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	if index, err := strconv.Atoi(indexstring); err == nil && index < len(todos) {
		c.JSON(http.StatusOK, gin.H{"item": todos[index]})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func AddListItem(c *gin.Context) {
	if verify(c) {
		item := c.PostForm("item")
		todos = append(todos, item)
		c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
	}
}

func verify(c *gin.Context) bool {
	status := true
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		verifierSetup := jwtverifier.JwtVerifier{
			Issuer:           "https://" + os.Getenv("OKTA_DOMAIN") + "/oauth2/default",
			ClaimsToValidate: toValidate,
		}
		verifier := verifierSetup.New()
		_, err := verifier.VerifyAccessToken(token)
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			print(err.Error())
			status = false
		}
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		status = false
	}
	return status
}

func main() {
	todos = append(todos, "Write the application")
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/lists", Lists)
	r.GET("/api/lists/:index", ListItem)
	r.POST("/api/lists", AddListItem)
	r.Run()
}
