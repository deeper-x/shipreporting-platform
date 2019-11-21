package webserver

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var inst = Instance{}

// // Most of webserver testing is delegated to integration testing, so %coverage here will not be huge
// func TestEngineBuild(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestAuth(t *testing.T) {
// 	resp, err := http.PostForm("http://127.0.0.1:8080/enter", url.Values{"username": {"xxxx"}, "password": {"xxxxx"}})

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	assert.Equal(t, http.StatusPermanentRedirect, resp.StatusCode)
// }

// func TestMultiRenderer(t *testing.T) {
// 	ok, _ := multiRenderer()

// 	assert.True(t, ok)
// }

// func TestMooringNow(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/mooring_now", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestAnchoredNow(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/anchored_now", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestArrivalsToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/arrivals_today", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestDeparturesToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/departures_today", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestArrivalPrevisionsNow(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/arrival_previsions_now", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestShippedGoodsToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/shipped_goods_today", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestTrafficListToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/traffic_list_today", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestShiftingPrevisionsToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/shifting_previsions_now", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }

// func TestDeparturePrevisionsToday(t *testing.T) {
// 	err := inst.EngineBuild()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	inst.Route()

// 	w := httptest.NewRecorder()

// 	req, err := http.NewRequest("GET", "/application/departure_previsions_now", nil)

// 	if err != nil {
// 		log.Println(err)
// 	}

// 	inst.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// }
