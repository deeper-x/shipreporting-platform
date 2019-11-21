package webserver

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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

// func TestLoginForm(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	err = godotenv.Load(utils.DotenvFile)

// 	data := url.Values{}
// 	data.Set("username", os.Getenv("TEST_USERNAME"))
// 	data.Set("password", os.Getenv("TEST_PASSWORD"))

// 	client := &http.Client{}

// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/enter", strings.NewReader(data.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }

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
	inst.LoadSession()

	err := inst.EngineBuild()

	inst.LoadSession()
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
