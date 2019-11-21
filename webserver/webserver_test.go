package webserver

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/deeper-x/shipreporting-platform/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var inst = Instance{}

// Most of webserver testing is delegated to integration testing, so %coverage here will not be huge
func TestEngineBuild(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}
}

func TestMultiRenderer(t *testing.T) {
	ok, _ := multiRenderer()

	assert.True(t, ok)
}

func TestMooringNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/mooring_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestLoginForm(t *testing.T) {
	err := godotenv.Load(utils.DotenvFile)

	credentials := url.Values{
		"username": {os.Getenv("TEST_USERNAME")},
		"password": {os.Getenv("TEST_PASSWORD")},
	}

	resp, err := http.PostForm("http://127.0.0.1:8080/enter", credentials)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAnchoredNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/anchored_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestArrivalsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/arrivals_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeparturesToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/departures_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestArrivalPrevisionsNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/arrival_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestShippedGoodsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/shipped_goods_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestTrafficListToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/traffic_list_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestShiftingPrevisionsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/shifting_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeparturePrevisionsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/application/departure_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
