package auth

import (
	"net/http/httptest"
	"testing"
)

func TestReadTokenAuth(t *testing.T) {
	rr := httptest.NewRecorder()

	mockResponse := rr.Result()
	msg, _ := ReadTokenAuth(mockResponse)

	if msg != "empty" {
		t.Fail()
	}
}
