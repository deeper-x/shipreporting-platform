package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/deeper-x/shipreporting-platform/utils"
)

// Credential response token data
type Credential struct {
	Token string `json:"token"`
}

// ReadTokenAuth get auth data from server
func ReadTokenAuth(username string, password string) (string, error) {
	
	err := godotenv.Load(utils.DotenvFile)

	if err != nil {
		log.Fatal(err)
	}

	cr := new(Credential)

	resp, err := http.PostForm(os.Getenv("GET_TOKEN_URL"), url.Values{"username": {username}, "password": {password}})

	if err != nil {
		log.Println(err)
		return "no data", err
	}

	defer resp.Body.Close()

	// req.Header.Set("Authorization", "Token ac772f9ebdb24b3b3399c30d82a63f56c808d070")

	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		log.Println(err)
		return "no data", err
	}

	return cr.Token, nil
}
