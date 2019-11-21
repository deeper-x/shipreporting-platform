package webserver

import (
	"log"
	"net/http"
	"strings"

	"github.com/deeper-x/shipreporting-platform/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var userkey = "shipreporting-session"
var token string
var pt *string

// AuthRequired middleware for restricted content
var AuthRequired = func(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Continue down the chain to handler etc
	c.Next()
}

// ProcessAuth auth processing
var ProcessAuth = func(c *gin.Context) (int, string) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	lenUsername := len(strings.Trim(username, " "))
	lenPassword := len(strings.Trim(password, " "))

	if lenUsername == 0 || lenPassword == 0 {
		return http.StatusBadRequest, "Empty credentials"
	}

	token, err := auth.ReadTokenAuth(username, password)

	pt = &token

	if err != nil {
		return http.StatusUnauthorized, "Credentials error - Not authorized"
	}

	if len(token) == 0 {
		return http.StatusUnauthorized, "Token error - Not authorized"
	}

	return http.StatusOK, "Authorized"
}

// CreateSession build session
var CreateSession = func(c *gin.Context) bool {
	session := sessions.Default(c)

	session.Set(userkey, *pt)

	if err := session.Save(); err != nil {
		log.Println(err)
		return false
	}

	return true
}

// DestroySession destroy session
var DestroySession = func(c *gin.Context) (int, string) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		return http.StatusBadRequest, "Invalid session token"
	}
	session.Delete(userkey)

	if err := session.Save(); err != nil {
		return http.StatusInternalServerError, "Failed to store session"
	}

	return http.StatusOK, "Successfully logged out"
}
