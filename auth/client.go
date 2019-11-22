package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/deeper-x/shipreporting-platform/utils"
	"github.com/joho/godotenv"
)

// Credential response token data
type Credential struct {
	Token string `json:"token"`
}

var cr = new(Credential)

// SignOn perform remote auth
func SignOn(username string, password string) (*http.Response, error) {
	err := godotenv.Load(utils.DotenvFile)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.PostForm(os.Getenv("GET_TOKEN_URL"), url.Values{"username": {username}, "password": {password}})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}

// ReadTokenAuth get auth data from server
func ReadTokenAuth(data *http.Response) (string, error) {

	if err := json.NewDecoder(data.Body).Decode(&cr); err != nil {
		log.Println(err)
		return "empty", err
	}

	defer data.Body.Close()

	return cr.Token, nil
}
