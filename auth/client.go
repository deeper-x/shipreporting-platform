package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

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

	// MORE CONCISE:
	// resp, err := http.PostForm(os.Getenv("GET_TOKEN_URL"), url.Values{"username": {username}, "password": {password}})

	// if err != nil {
	// 	return nil, err
	// }

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	client := &http.Client{}

	req, err := http.NewRequest("POST", os.Getenv("GET_TOKEN_URL"), strings.NewReader(data.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	return resp, err
}

// ReadTokenAuth get auth data from server
func ReadTokenAuth(data *http.Response) (string, error) {
	if err := json.NewDecoder(data.Body).Decode(&cr); err != nil {
		return "", err
	}

	defer data.Body.Close()

	return cr.Token, nil
}
