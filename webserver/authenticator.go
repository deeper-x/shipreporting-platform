package webserver

import (
	"log"
	"net/http"
	"strings"

	"github.com/deeper-x/shipreporting-platform/auth"
	"github.com/deeper-x/shipreporting-platform/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var sessionKey = "shipreporting-session"
var portinformer = "portinformer"
var token string
var ptoken *string

var connector = db.Connector()
var repository = db.Repository{Conn: connector}

// AuthRequired middleware for restricted content
var AuthRequired = func(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(sessionKey)

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	return

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

	resp, err := auth.SignOn(username, password)

	if err != nil {
		log.Println(err)
	}

	token, err := auth.ReadTokenAuth(resp)

	ptoken = &token

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

	userID := GetUserID(repository, *ptoken)
	portinformer := GetManagedPortinformer(repository, userID)

	session.Set(portinformer, userID)
	session.Set(sessionKey, *ptoken)
	session.Set("managedPortinformer", portinformer)

	repository.Close()

	if err := session.Save(); err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GetUserID retrive user id from auth token
var GetUserID = func(repo db.Repository, authToken string) string {
	return repo.SelectUserID(authToken)
}

// GetManagedPortinformer retrieve managed portinformer
var GetManagedPortinformer = func(repo db.Repository, userID string) string {
	return repo.SelectUserPortinformer(userID)
}

// DestroySession destroy session
var DestroySession = func(c *gin.Context) (int, string) {
	session := sessions.Default(c)
	user := session.Get(sessionKey)

	if user == nil {
		return http.StatusBadRequest, "Invalid session token"
	}
	session.Delete(sessionKey)

	if err := session.Save(); err != nil {
		return http.StatusInternalServerError, "Failed to store session"
	}

	return http.StatusOK, "Successfully logged out"
}
