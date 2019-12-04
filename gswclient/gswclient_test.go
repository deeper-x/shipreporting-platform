package gswclient

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchData(t *testing.T) {
	_, err := FetchData("http://80.88.88.162:8000/moored/27")

	if err != nil {
		log.Println(err)
		assert.Fail(t, "GSW client error")
	}

}

func TestJsonConvert(t *testing.T) {
	_, err := JSONConvert([]byte("[{\"example\": \"content\"}]"))

	if err != nil {
		log.Println(err)
		assert.Fail(t, "JSON conversion failed")
	}
}
